package main

import (
	"flag"
	"fmt"
	pbManagement "go-sdap/src/proto/management"
	pbSdap "go-sdap/src/proto/sdap"
	"go-sdap/src/server/dbManager"
	"go-sdap/src/server/managementServer"
	"go-sdap/src/server/operationServer"
	"log/slog"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	operationPort  = flag.Int("operationPort", 50051, "The operation server port")
	managementPort = flag.Int("managementPort", 50052, "The management server port")
)

func checkDatabaseConnection(logger *slog.Logger, db *dbManager.DbManager) {
	for {
		err := db.Ping()
		if err != nil {
			logger.Error("failed to ping database, reconnecting...", "error", err)

			db.Reconnect()
		}

		time.Sleep(5 * time.Second)
	}
}

func startOperationServer(logger *slog.Logger, db *dbManager.DbManager) {
	for {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *operationPort))
		if err != nil {
			logger.Error("failed to listen", "error", err)
			time.Sleep(2 * time.Second)
			continue
		}
		s := grpc.NewServer()
		pbSdap.RegisterOperationServer(s, operationServer.New(logger, db))
		logger.Info("Operation server listening at", "address", lis.Addr())
		if err := s.Serve(lis); err != nil {
			logger.Error("failed to serve", "error", err)
			s.GracefulStop()
		}
	}
}

func startManagementServer(logger *slog.Logger, db *dbManager.DbManager) {
	for {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *managementPort))
		if err != nil {
			logger.Error("failed to listen", "error", err)
			return
		}
		s := grpc.NewServer()
		pbManagement.RegisterManagementServer(s, managementServer.New(logger, db))
		logger.Info("Management server listening at", "address", lis.Addr())
		if err := s.Serve(lis); err != nil {
			logger.Error("failed to serve", "error", err)
			s.GracefulStop()
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	flag.Parse()
	logger := slog.Default()

	db := dbManager.New(logger)
	defer db.Disconnect()

	go checkDatabaseConnection(logger, db)
	go startOperationServer(logger, db)
	go startManagementServer(logger, db)

	select {}
}
