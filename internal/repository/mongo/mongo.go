package mongo

import (
	"context"

	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client            *mongo.Client
	productCollection *mongo.Collection
)

func PingMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return Client.Ping(ctx, nil)
}

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
	productCollection = Client.Database("ecommerce").Collection("products")
}

// TODO Pagination?
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

func FetchProductByID(id string) (models.Product, error) {
	var product models.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err 
	}

	err = productCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
	return product, err
}

func InsertProduct(product models.Product) error {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := productCollection.InsertOne(context.TODO(), product)
	return err
}

func UpdateProduct(id string, product models.Product) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now()) 

	_, err = productCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": product},
	)
	return err
}

func RemoveProduct(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = productCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
