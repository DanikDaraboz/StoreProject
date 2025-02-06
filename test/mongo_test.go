package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoConnection(t *testing.T) {
	mongoURI := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(ctx)

	fmt.Println("Connected to MongoDB")

	// Get database & collection
	db := client.Database("ecommerce")
	collection := db.Collection("products")

	// Insert a test product
	_, err = collection.InsertOne(ctx, bson.M{"name": "Test Product", "price": 100})
	if err != nil {
		t.Fatalf("Failed to insert document: %v", err)
	}
	fmt.Println("Inserted test product")

	// Find the inserted product
	var result bson.M
	err = collection.FindOne(ctx, bson.M{"name": "Test Product"}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find document: %v", err)
	}
	fmt.Println("Retrieved document:", result)
}
