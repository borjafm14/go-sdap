package main

import (
	"flag"
	"fmt"
	pbManagement "go-sdap/src/proto/management"
	pbSdap "go-sdap/src/proto/sdap"
	"go-sdap/src/server/dbManager"
	"go-sdap/src/server/managementServer"
	"go-sdap/src/server/operationServer"
	"io"
	"log/slog"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

var (
	operationPort  = flag.Int("operationPort", 50051, "The operation server port")
	managementPort = flag.Int("managementPort", 50052, "The management server port")
	logDir         = flag.String("logDir", "log", "The directory to store logs")
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

func rotateLogs(logDir string, logger *slog.Logger, logFile *os.File) {
	var err error
	for {
		time.Sleep(24 * time.Hour)
		logFile.Close()
		os.Rename(fmt.Sprintf("%s/sdap.log", logDir), fmt.Sprintf("%s/sdap-%s.log", logDir, time.Now().Format("2006-01-02")))
		logFile, err = os.OpenFile(fmt.Sprintf("%s/sdap.log", logDir), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Failed to open log file: %v\n", err)
			return
		}
		multiWriter := io.MultiWriter(logFile, os.Stdout)
		*logger = *slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}

func main() {
	flag.Parse()
	if err := os.MkdirAll(*logDir, 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
		return
	}

	logFile, err := os.OpenFile(fmt.Sprintf("%s/sdap.log", *logDir), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	logger := slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: slog.LevelInfo}))

	db := dbManager.New(logger)
	defer db.Disconnect()

	go rotateLogs(*logDir, logger, logFile)
	go checkDatabaseConnection(logger, db)
	go startOperationServer(logger, db)
	go startManagementServer(logger, db)

	select {}
}
