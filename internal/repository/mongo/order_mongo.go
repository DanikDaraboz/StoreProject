package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrders() ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database("ecommerce").Collection("orders")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func FetchOrderByID(id string) (models.Order, error) {
	var order models.Order
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	err = orderCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&order)
	return order, err
}

func InsertOrder(order models.Order) (primitive.ObjectID, error) {
	order.ID = primitive.NewObjectID()
	order.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	order.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to get inserted ID")
	}

	return insertedID, nil
}

func UpdateOrder(id string, order models.Order) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	order.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err = orderCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": order},
	)
	return err
}

func RemoveOrder(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = orderCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
