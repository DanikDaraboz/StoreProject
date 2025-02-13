package mongo

import (
	"context"
	"time"

	"github.com/DanikDaraboz/StoreProject/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client            *mongo.Client
	productCollection *mongo.Collection
	orderCollection   *mongo.Collection
)

func PingMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return Client.Ping(ctx, nil)
}

func Connect(mongoURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logger.ErrorLogger.Println("Failed to create MongoDB client:", err)
		return nil, err 
	}

	err = client.Connect(ctx)
	if err != nil {
		logger.ErrorLogger.Println("Failed to connect to MongoDB:", err)
		return nil, err 
	}

	logger.InfoLogger.Println("Connected to MongoDB at", mongoURI)
	
	Client = client
	productCollection = Client.Database("ecommerce").Collection("products")
	orderCollection = Client.Database("ecommerce").Collection("orders")

	return client, nil 
}
