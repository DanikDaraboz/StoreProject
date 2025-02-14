package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
)

var _ interfaces.SessionServicesInterface = (*sessionServices)(nil)

type sessionServices struct {
	sessionRepo repoInterface.SessionRepositoryInterface
}

func NewSessionServices(sessionRepo repoInterface.SessionRepositoryInterface) interfaces.SessionServicesInterface {
	return &sessionServices{sessionRepo: sessionRepo}
}

func (s sessionServices) CreateSession(userID string) (string, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}
	expiresAt := time.Now().Add(24 * time.Hour)

	err = s.sessionRepo.InsertSession(sessionID, userID, expiresAt)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func (s sessionServices) FindSession(sessionID string) (string, error) {
	session, err := s.sessionRepo.FindSessionByID(sessionID)
	if err != nil {
		return "", err
	}

	if session.ExpiresAt.Before(time.Now()) {
		err = s.sessionRepo.DeleteSessionByID(sessionID)
		if err != nil {
			return "", err
		}
		return "", errors.New("session expired")
	}

	return session.UserID, nil
}

func (s sessionServices) DeleteSession(sessionID string) error {
	err := s.sessionRepo.DeleteSessionByID(sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (s sessionServices) ClearExpiredSessions() error {
	err := s.sessionRepo.DeleteExpiredSessions()
	if err != nil {
		return err
	}

	return nil
}

// Helper to generate sessionID
func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", errors.New("failed to generate session ID")
	}
	return hex.EncodeToString(bytes), nil
}
