package managementServer

import (
	pb "go-sdap/src/proto/sdap"
	"log/slog"
)

type managementServer struct {
	pb.UnimplementedManagementServer
	logger *slog.Logger
}

func New(logger *slog.Logger) *managementServer {
	return &managementServer{
		logger: logger,
	}
}
