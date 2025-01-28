package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	ExpiresAt time.Time `bson:"expires_at"`
}

// Helper to generate sessionID
func generateSessionID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", errors.New("failed to generate session ID")
	}
	return hex.EncodeToString(bytes), nil
}

func CreateSession(client *mongo.Client, userID string) (string, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}

	session := Session{
		ID:        sessionID,
		UserID:    userID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	collection := client.Database("store").Collection("sessions")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, session)
	if err != nil {
		return "", errors.New("failed to create session")
	}

	log.Println("Session created:", sessionID, "for user:", userID)
	return sessionID, nil
}

func ValidateSession(client *mongo.Client, sessionID string) (string, error) {
	collection := client.Database("store").Collection("sessions")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session Session
	err := collection.FindOne(ctx, bson.M{
		"_id":        sessionID,
		"expires_at": bson.M{"$gt": time.Now()},
	}).Decode(&session)
	if err != nil {
		return "", errors.New("invalid session")
	}

	return session.UserID, nil
}

func DeleteSession(client *mongo.Client, sessionID string) error {
	collection := client.Database("store").Collection("sessions")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": sessionID})
	if err != nil {
		return errors.New("failed to delete session")
	}

	log.Println("Session deleted:", sessionID)
	return nil
}
