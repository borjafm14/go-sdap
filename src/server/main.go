package main

import (
	"flag"
	"fmt"
	pb "go-sdap/src/proto/sdap"
	"go-sdap/src/server/operationServer"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

var (
	operationPort  = flag.Int("operationPort", 50051, "The operation server port")
	managementPort = flag.Int("managementPort", 50052, "The management server port")
)

func main() {
	flag.Parse()
	logger := slog.Default()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *operationPort))
	if err != nil {
		logger.Error("failed to listen", "error", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterOperationServer(s, operationServer.New(logger))
	logger.Info("server listening at", "address", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve", "error", err)
	}
}
