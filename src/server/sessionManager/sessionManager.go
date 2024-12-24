package sessionManager

import (
	"log/slog"
	"sync"
	"time"
)

type Session struct {
	Username      string
	Token         string
	Authenticated bool
	Timestamp     time.Time
}

type SessionManager struct {
	logger   *slog.Logger
	sessions map[string]*Session
	mu       sync.Mutex
}

func New(logger *slog.Logger) *SessionManager {
	sm := &SessionManager{
		logger:   logger,
		sessions: make(map[string]*Session),
	}

	go func() {
		for {
			time.Sleep(1 * time.Minute)
			sm.CleanupExpiredSessions()
		}
	}()

	return sm
}

func (sm *SessionManager) CreateSession(username, token string, authenticated bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.sessions[token] = &Session{
		Username:      username,
		Token:         token,
		Authenticated: authenticated,
		Timestamp:     time.Now(),
	}
	sm.logger.Info("Session created", "username", username, "token", token)
}

func (sm *SessionManager) GetSession(token string) (*Session, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	session, exists := sm.sessions[token]
	return session, exists
}

func (sm *SessionManager) UpdateSessionTimestamp(token string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if session, exists := sm.sessions[token]; exists {
		session.Timestamp = time.Now()
		sm.logger.Info("Session timestamp updated", "token", token)
	}
}

func (sm *SessionManager) CleanupExpiredSessions() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	for token, session := range sm.sessions {
		if time.Since(session.Timestamp) > 10*time.Minute {
			delete(sm.sessions, token)
			sm.logger.Info("Session expired and removed", "token", token)
		}
	}
}
