package managementServer

import (
	"context"
	pb "go-sdap/src/proto/management"
	"go-sdap/src/server/dbManager"
	"go-sdap/src/server/helper"
	"go-sdap/src/server/sessionManager"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type managementServer struct {
	pb.UnimplementedManagementServer
	logger *slog.Logger
	db     *dbManager.DbManager
	sm     *sessionManager.SessionManager
}

func New(logger *slog.Logger, db *dbManager.DbManager, sm *sessionManager.SessionManager) *managementServer {

	if db != nil {
		// Create admin user if it doesn't exist
		db.CreateAdminUser()
	}

	return &managementServer{
		logger: logger,
		db:     db,
		sm:     sm,
	}
}

func (s *managementServer) Connect(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
	logger := s.logger.With("RPC", "Connect")
	logger.Info("Incoming request", "req", in)

	status := pb.Status_STATUS_OK
	token, err := s.sm.CreateSession(in.Hostname)

	if err != nil {
		status = pb.Status_STATUS_ERROR
	}

	return &pb.SessionResponse{
		Token:  token,
		Status: status,
	}, nil
}

func (s *managementServer) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	logger := s.logger.With("RPC", "Authenticate")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) {
		return &pb.AuthenticateResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if s.sm.IsAuthenticated(in.Token) {
		return &pb.AuthenticateResponse{
			Status: pb.Status_STATUS_OK,
		}, nil
	}

	if s.db == nil {
		return &pb.AuthenticateResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status := s.db.AuthenticateAdmin(in.Username, in.Password)

	if status == pb.Status_STATUS_OK {
		s.sm.SetAuthenticated(in.Token, in.Username)
	}

	return &pb.AuthenticateResponse{
		Status: status,
	}, nil
}

func (s *managementServer) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	logger := s.logger.With("RPC", "GetUser")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.UserResponse{
			User:   nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if in.Username == "admin" {
		logger.Warn("Admin user cannot be retrieved")
		return &pb.UserResponse{
			User:   nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.UserResponse{
			User:   nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status, user := s.db.GetUser(in.Username)

	return &pb.UserResponse{
		User:   user,
		Status: status,
	}, nil
}

func (s *managementServer) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	logger := s.logger.With("RPC", "ListUsers")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.ListUsersResponse{
			Users:  nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if in.Username == helper.StringPtr("admin") {
		logger.Warn("Admin user cannot be listed")
		return &pb.ListUsersResponse{
			Users:  nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.ListUsersResponse{
			Users:  nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	users, status := s.db.ListUsers(in.Username, in.Filter)

	return &pb.ListUsersResponse{
		Users:  users,
		Status: status,
	}, nil
}

func (s *managementServer) ModifyUsers(ctx context.Context, in *pb.ModifyUsersRequest) (*pb.ModifyUsersResponse, error) {
	logger := s.logger.With("RPC", "ModifyUsers")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.ModifyUsersResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if len(in.Usernames) > 0 {
		for _, username := range in.Usernames {
			if username == "admin" {
				logger.Warn("Admin user cannot be modified")
				return &pb.ModifyUsersResponse{
					Status: pb.Status_STATUS_ERROR,
				}, nil
			}
		}
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.ModifyUsersResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status := s.db.ModifyUsers(in.Usernames, in.Filter)

	return &pb.ModifyUsersResponse{
		Status: status,
	}, nil
}

func (s *managementServer) ChangeUsername(ctx context.Context, in *pb.UsernameRequest) (*pb.UsernameResponse, error) {
	logger := s.logger.With("RPC", "ChangeUsername")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.UsernameResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if in.OldUsername == "admin" || in.NewUsername == "admin" {
		logger.Warn("Admin user cannot be modified")
		return &pb.UsernameResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.UsernameResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status := s.db.ChangeUsername(in.OldUsername, in.NewUsername)

	return &pb.UsernameResponse{
		Status: status,
	}, nil
}

func (s *managementServer) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	logger := s.logger.With("RPC", "ChangePassword")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.ChangePasswordResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if in.Username != "admin" {
		logger.Warn("Only admin's password can be changed")
		return &pb.ChangePasswordResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.ChangePasswordResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status := s.db.ChangeAdminPassword(in.OldPassword, in.NewPassword)

	return &pb.ChangePasswordResponse{
		Status: status,
	}, nil
}

func (s *managementServer) AddUsers(ctx context.Context, in *pb.AddUsersRequest) (*pb.AddUsersResponse, error) {
	logger := s.logger.With("RPC", "AddUsers")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.AddUsersResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if len(in.Users) > 0 {
		for _, user := range in.Users {
			if user.Username == helper.StringPtr("admin") {
				logger.Warn("Admin user cannot be added")
				return &pb.AddUsersResponse{
					Status: pb.Status_STATUS_ERROR,
				}, nil
			}
		}
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.AddUsersResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status := s.db.AddUsers(in.Users)

	return &pb.AddUsersResponse{
		Status: status,
	}, nil
}

func (s *managementServer) DeleteUsers(ctx context.Context, in *pb.DeleteUsersRequest) (*pb.DeleteUsersResponse, error) {
	logger := s.logger.With("RPC", "DeleteUsers")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.DeleteUsersResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if len(in.Usernames) > 0 {
		for _, username := range in.Usernames {
			if username == "admin" {
				logger.Warn("Admin user cannot be deleted")
				return &pb.DeleteUsersResponse{
					Status: pb.Status_STATUS_ERROR,
				}, nil
			}
		}
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.DeleteUsersResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	status := s.db.DeleteUsers(in.Usernames)

	return &pb.DeleteUsersResponse{
		Status: status,
	}, nil
}

func (s *managementServer) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*emptypb.Empty, error) {
	logger := s.logger.With("RPC", "Disconnect")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &emptypb.Empty{}, nil
	}

	s.sm.DeleteSession(in.Token)

	return &emptypb.Empty{}, nil
}
