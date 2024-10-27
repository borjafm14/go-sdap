package operationServer

import (
	pb "go-sdap/src/proto/sdap"
	"log/slog"
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
