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

var _ interfaces.CartRepositoryInterface = (*cartRepository)(nil)

type cartRepository struct {
	collection *mongo.Collection
}

func NewCartRepository(collection *mongo.Collection) interfaces.CartRepositoryInterface {
	return &cartRepository{collection: collection}
}

func (c *cartRepository) InsertCart(cart *models.Cart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.collection.InsertOne(ctx, cart)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
		return err
	}

	return nil
}

func (c *cartRepository) FindCartByUserID(userID primitive.ObjectID) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cart models.Cart
	// Query by "user_id" instead of "_id"
	err := c.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
	}
	return &cart, err
}

func (c *cartRepository) UpdateCart(cart *models.Cart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare the update document
	update := bson.M{
		"$set": bson.M{
			"items":       cart.Items,
			"total_price": cart.TotalPrice,
		},
	}

	result, err := c.collection.UpdateOne(ctx, bson.M{"user_id": cart.UserID}, update)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
		return err
	}

	if result.MatchedCount == 0 {
		err = errors.New("no cart found to update")
		logger.ErrorLogger.Println("error", err)
		return err
	}

	return nil
}
