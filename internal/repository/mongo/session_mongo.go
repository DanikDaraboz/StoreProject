package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.SessionRepositoryInterface = (*sessionRepo)(nil)

type sessionRepo struct {
	collection *mongo.Collection
}

func NewSessionRepository(collection *mongo.Collection) interfaces.SessionRepositoryInterface {
	return &sessionRepo{collection: collection}
}

func (s *sessionRepo) InsertSession(sessionID string, userID primitive.ObjectID, expiresAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// No need to convert to ObjectID if your struct uses strings
	_, err := s.collection.InsertOne(ctx, bson.M{
		"_id":        sessionID,
		"user_id":    userID,
		"expires_at": expiresAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionRepo) FindSessionByID(sessionID string) (*models.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session models.Session
	err := s.collection.FindOne(ctx, bson.M{"_id": sessionID}).Decode(&session)
	if err != nil {
		return &models.Session{}, err
	}

	return &session, nil
}

func (s *sessionRepo) DeleteSessionByID(sessionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": sessionID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no session found to delete")
	}

	return nil
}

func (s *sessionRepo) DeleteExpiredSessions() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filter expired sessions
	filter := bson.M{"expires_at": bson.M{"$lt": time.Now()}}

	result, err := s.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount > 0 {
		logger.ErrorLogger.Printf("%d expired sessions deleted\n", result.DeletedCount)
	}

	return nil
}
