package services

import (
	"errors"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.SessionServicesInterface = (*sessionServices)(nil)

type sessionServices struct {
	sessionRepo repoInterface.SessionRepositoryInterface
}

func NewSessionServices(sessionRepo repoInterface.SessionRepositoryInterface) interfaces.SessionServicesInterface {
	return &sessionServices{sessionRepo: sessionRepo}
}

func (s sessionServices) CreateSession(userID primitive.ObjectID) (string, error) {
	sessionID, err := utils.GenerateSessionID()
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

func (s sessionServices) FindSession(sessionID string) (models.Session, error) {
	session, err := s.sessionRepo.FindSessionByID(sessionID)
	if err != nil {
		return models.Session{}, err
	}

	if session.ExpiresAt.Before(time.Now()) {
		err = s.sessionRepo.DeleteSessionByID(sessionID)
		if err != nil {
			return models.Session{}, err
		}
		return models.Session{}, errors.New("session expired")
	}

	return session, nil
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
