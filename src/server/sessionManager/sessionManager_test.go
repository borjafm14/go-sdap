package sessionManager

import (
	"os"
	"testing"
	"time"

	"log/slog"
)

func TestSessionManager_CreateSession(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	sm := New(logger)

	hostname := "test-host"
	token, err := sm.CreateSession(hostname)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	session, exists := sm.GetSession(token)
	if !exists {
		t.Fatalf("expected session to exist")
	}

	if session.Hostname != hostname {
		t.Errorf("expected hostname %v, got %v", hostname, session.Hostname)
	}
}

func TestSessionManager_GetSession(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	session, exists := sm.GetSession(token)
	if !exists {
		t.Fatalf("expected session to exist")
	}

	if session.Hostname != hostname {
		t.Errorf("expected hostname %v, got %v", hostname, session.Hostname)
	}
}

func TestSessionManager_SessionExists(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	if !sm.SessionExists(token) {
		t.Fatalf("expected session to exist")
	}
}

func TestSessionManager_IsAuthenticated(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	if sm.IsAuthenticated(token) {
		t.Fatalf("expected session to be unauthenticated")
	}

	username := "test-user"
	sm.SetAuthenticated(token, username)

	if !sm.IsAuthenticated(token) {
		t.Fatalf("expected session to be authenticated")
	}
}

func TestSessionManager_SetAuthenticated(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	username := "test-user"
	sm.SetAuthenticated(token, username)

	session, _ := sm.GetSession(token)
	if session.Username == nil || *session.Username != username {
		t.Errorf("expected username %v, got %v", username, session.Username)
	}

	if !session.Authenticated {
		t.Errorf("expected session to be authenticated")
	}
}

func TestSessionManager_DeleteSession(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	sm.DeleteSession(token)

	if sm.SessionExists(token) {
		t.Fatalf("expected session to be deleted")
	}
}

func TestSessionManager_UpdateSessionTimestamp(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	oldTimestamp := sm.sessions[token].Timestamp
	time.Sleep(1 * time.Second)
	sm.UpdateSessionTimestamp(token)

	if sm.sessions[token].Timestamp == oldTimestamp {
		t.Errorf("expected timestamp to be updated")
	}
}

func TestSessionManager_CleanupExpiredSessions(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	sm := New(logger)

	hostname := "test-host"
	token, _ := sm.CreateSession(hostname)

	sm.sessions[token].Timestamp = time.Now().Add(-11 * time.Minute)
	sm.CleanupExpiredSessions()

	if sm.SessionExists(token) {
		t.Fatalf("expected session to be cleaned up")
	}
}
