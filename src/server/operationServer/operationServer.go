package operationServer

import (
	"context"
	pb "go-sdap/src/proto/sdap"
	"go-sdap/src/server/dbManager"
	"go-sdap/src/server/sessionManager"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type operationServer struct {
	pb.UnimplementedOperationServer
	logger *slog.Logger
	db     *dbManager.DbManager
	sm     *sessionManager.SessionManager
}

func New(logger *slog.Logger, db *dbManager.DbManager, sm *sessionManager.SessionManager) *operationServer {
	return &operationServer{
		logger: logger,
		db:     db,
		sm:     sm,
	}
}

func (s *operationServer) Connect(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
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

func (s *operationServer) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	logger := s.logger.With("RPC", "Authenticate")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) {
		return &pb.AuthenticateResponse{
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	if s.sm.IsAuthenticated(in.Token) {
		logger.Info("User already authenticated", "username", in.Username)
		return &pb.AuthenticateResponse{
			Status: pb.Status_STATUS_OK,
		}, nil
	}

	if s.db == nil {
		return &pb.AuthenticateResponse{
			User:   nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	user, status := s.db.Authenticate(in.Username, in.Password)

	return &pb.AuthenticateResponse{
		User:   user,
		Status: status,
	}, nil
}

func (s *operationServer) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	logger := s.logger.With("RPC", "ChangePassword")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
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

	status := s.db.ChangePassword(in.Username, in.OldPassword, in.NewPassword)

	return &pb.ChangePasswordResponse{
		Status: status,
	}, nil
}

func (s *operationServer) GetCharacteristics(ctx context.Context, in *pb.CharacteristicsRequest) (*pb.CharacteristicsResponse, error) {
	logger := s.logger.With("RPC", "GetCharacteristics")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.CharacteristicsResponse{
			User:   nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.CharacteristicsResponse{
			User:   nil,
			Status: pb.Status_STATUS_ERROR,
		}, nil
	}

	user, status := s.db.GetCharacteristics(in.Username, in.Characteristics)

	return &pb.CharacteristicsResponse{
		User:   user,
		Status: status,
	}, nil
}

func (s *operationServer) GetMemberOf(ctx context.Context, in *pb.MemberOfRequest) (*pb.MemberOfResponse, error) {
	logger := s.logger.With("RPC", "GetMemberOf")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &pb.MemberOfResponse{
			MemberOf: nil,
			Status:   pb.Status_STATUS_ERROR,
		}, nil
	}

	s.sm.UpdateSessionTimestamp(in.Token)

	if s.db == nil {
		return &pb.MemberOfResponse{
			MemberOf: nil,
			Status:   pb.Status_STATUS_ERROR,
		}, nil
	}

	memberOf, status := s.db.GetMemberOf(in.Username)

	return &pb.MemberOfResponse{
		MemberOf: memberOf,
		Status:   status,
	}, nil
}

func (s *operationServer) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*emptypb.Empty, error) {
	logger := s.logger.With("RPC", "Disconnect")
	logger.Info("Incoming request", "req", in)

	if !s.sm.SessionExists(in.Token) || !s.sm.IsAuthenticated(in.Token) {
		return &emptypb.Empty{}, nil
	}

	s.sm.DeleteSession(in.Token)

	return &emptypb.Empty{}, nil
}
