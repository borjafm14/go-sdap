package main

import (
	"go-sdap/src/client/managementClient"
	"go-sdap/src/client/sdapClient"
	"log/slog"
)

func main() {
	logger := slog.Default()

	logger.Info("Creating SDAP client...")

	s := sdapClient.New()
	status, err := s.Connect("127.0.0.1", 50051, false)
	if err == nil {
		logger.Info("Connect", "status", status.String())
	}

	user, status, err := s.Authenticate("borja", "1234")
	if err == nil {
		logger.Info("Authenticate", "status", status.String())
		logger.Info("Authenticate", "user", user.String())
	}

	s.Disconnect()

	logger.Info("Creating management client...")
	m := managementClient.New()

	mstatus, err := m.Connect("127.0.0.1", 50052, false)
	if err == nil {
		logger.Info("Connect", "status", mstatus.String())
	}

	muser, mstatus, err := m.GetUser("borja")
	if err == nil {
		logger.Debug("GetUser", "status", mstatus.String())
		logger.Debug("GetUser", "user", muser.String())
	}

	m.Disconnect()
}
