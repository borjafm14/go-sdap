package operationServer

import (
	"context"
	pb "go-sdap/src/proto/sdap"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type operationServer struct {
	pb.UnimplementedOperationServer
	logger *slog.Logger
}

func New(logger *slog.Logger) *operationServer {
	return &operationServer{
		logger: logger,
	}
}

func (s *operationServer) Connect(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
	logger := s.logger.With("RPC", "Connect")
	logger.Info("Incoming request", "req", in)

	return &pb.SessionResponse{
		Token:  "1234",
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *operationServer) Authenticate(ctx context.Context, in *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	logger := s.logger.With("RPC", "Authenticate")
	logger.Info("Incoming request", "req", in)

	var user *pb.User
	return &pb.AuthenticateResponse{
		User:   user,
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *operationServer) GetCharacteristics(ctx context.Context, in *pb.CharacteristicsRequest) (*pb.CharacteristicsResponse, error) {
	logger := s.logger.With("RPC", "GetCharacteristics")
	logger.Info("Incoming request", "req", in)

	var user *pb.User
	return &pb.CharacteristicsResponse{
		User:   user,
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *operationServer) GetMemberOf(ctx context.Context, in *pb.MemberOfRequest) (*pb.MemberOfResponse, error) {
	logger := s.logger.With("RPC", "GetMemberOf")
	logger.Info("Incoming request", "req", in)

	var memberOf []string
	return &pb.MemberOfResponse{
		MemberOf: memberOf,
		Status:   pb.Status_STATUS_OK,
	}, nil
}

func (s *operationServer) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*emptypb.Empty, error) {
	logger := s.logger.With("RPC", "Disconnect")
	logger.Info("Incoming request", "req", in)

	return &emptypb.Empty{}, nil
}
