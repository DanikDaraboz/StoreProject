package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.OrderRepositoryInterface = (*orderRepository)(nil)

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) interfaces.OrderRepositoryInterface {
	return &orderRepository{collection: collection}
}

// TODO cursor.Next() instead of cursor.All()
func (o *orderRepository) GetOrders() ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query all orders
	cursor, err := o.collection.Find(ctx, bson.M{})
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

func (o *orderRepository) FetchOrderByID(id string) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var order models.Order
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	err = o.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	return order, err
}

func (o *orderRepository) InsertOrder(order models.Order) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order.ID = primitive.NewObjectID()
	order.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	order.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to get inserted ID")
	}

	return insertedID, nil
}

func (o *orderRepository) UpdateOrder(id string, order models.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	order.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err = o.collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": order},
	)
	return err
}

func (o *orderRepository) RemoveOrder(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = o.collection.DeleteOne(ctx, bson.M{"_id": objID})

	return err
}
