package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get database & collection
	db := client.Database("ecommerce")
	collection := db.Collection("products")

	// Insert 
	_, err = collection.InsertOne(ctx, bson.M{"name": "Test Product", "price": 100})
	if err != nil {
		log.Fatal("Failed to insert document:", err)
	}
	fmt.Println("Inserted test product!")

	// Find 
	var result bson.M
	err = collection.FindOne(ctx, bson.M{"name": "Test Product"}).Decode(&result)
	if err != nil {
		log.Fatal("Failed to find document:", err)
	}
	fmt.Println("Retrieved document:", result)
}
