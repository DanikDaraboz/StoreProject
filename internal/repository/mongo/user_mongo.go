package mongo

import (
	"context"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.UserRepositoryInterface = (*userRepository)(nil)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) interfaces.UserRepositoryInterface {
	return &userRepository{collection: collection}
}

func (u userRepository) InsertUser(user models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return objectID, nil
}

func (u userRepository) FindUserByID(userID primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := u.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)

	return user, err
}

func (u userRepository) FindUserByEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := u.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	return user, err
}

func (u userRepository) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"email":    user.Email,
			"age":      user.Age,
			"phone":    user.Phone,
			"address":  user.Address,
		},
	}

	_, err := u.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, update)
	if err != nil {
		return err
	}

	return nil
}
