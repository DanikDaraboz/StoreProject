package mongo

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
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
	return primitive.NilObjectID, nil
}

func (u userRepository) FindUserByID(id primitive.ObjectID) (models.User, error) {
	return models.User{}, nil
}

func (u userRepository) FindUserByEmail(email string) (models.User, error) {
	return models.User{}, nil
}

func (u userRepository) UpdateUser(user models.User) error {
	return nil
}

func (u userRepository) ManageSession(userID primitive.ObjectID, sessionKey string, action string) error{
	return nil
}
