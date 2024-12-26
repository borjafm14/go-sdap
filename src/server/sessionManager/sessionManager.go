package sessionManager

import (
	"log/slog"
	"sync"
	"time"

	"go-sdap/src/server/helper"
)

type Session struct {
	Hostname      string
	Username      *string
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

func (sm *SessionManager) CreateSession(hostname string) (string, error) {
	token := helper.GenerateToken()

	// check that token is unique
	for {
		sm.mu.Lock()
		_, exists := sm.sessions[token]
		sm.mu.Unlock()

		if exists {
			token = helper.GenerateToken()
		} else {
			break
		}
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.sessions[token] = &Session{
		Hostname:      hostname,
		Username:      nil,
		Token:         token,
		Authenticated: false,
		Timestamp:     time.Now(),
	}
	sm.logger.Info("Session created", "hostname", hostname, "token", token)

	return token, nil
}

func (sm *SessionManager) GetSession(token string) (*Session, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	session, exists := sm.sessions[token]
	return session, exists
}

func (sm *SessionManager) SessionExists(token string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	_, exists := sm.sessions[token]
	return exists
}

func (sm *SessionManager) IsAuthenticated(token string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if session, exists := sm.sessions[token]; exists {
		return session.Username != nil && session.Authenticated
	}
	return false
}

func (sm *SessionManager) SetAuthenticated(token string, username string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if session, exists := sm.sessions[token]; exists {
		session.Username = &username
		session.Authenticated = true
		session.Timestamp = time.Now()
		sm.logger.Info("Session authenticated", "token", token, "username", username)
	}
}

func (sm *SessionManager) DeleteSession(token string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, token)
	sm.logger.Info("Session deleted", "token", token)
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
