package mongo

import (
	"context"

	"time"

	"github.com/DanikDaraboz/StoreProject/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(mongoURI string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logger.ErrorLogger.Println("Failed to create MongoDB client:", err)
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		logger.ErrorLogger.Println("Failed to connect to MongoDB:", err)
		return
	}

	logger.InfoLogger.Println("Connected to MongoDB at", mongoURI)
	Client = client
}

func GetProducts() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database("ecommerce").Collection("products")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []map[string]interface{}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func PingMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return Client.Ping(ctx, nil)
}
