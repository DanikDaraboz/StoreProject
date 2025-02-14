package services

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.UserServicesInterface = (*userServices)(nil)

type userServices struct {
	userRepo repoInterface.UserRepositoryInterface
}

func NewUserServices(userRepo repoInterface.UserRepositoryInterface) interfaces.UserServicesInterface {
	return &userServices{userRepo: userRepo}
}

func (u userServices) RegisterUser(user models.User) (primitive.ObjectID, error) {
	return primitive.NilObjectID, nil
}

// Returns session key
func (u userServices) LoginUser(email string, password string) (string, error) {
	return "", nil
}

func (u userServices) LogoutUser(sessionKey string) error {
	return nil
}

func (u userServices) GetUser(id string) (models.User, error) {
	return models.User{}, nil
}

func (u userServices) UpdateUser(user models.User) error {
	return nil
}

func (u userServices) ManageSession(userID string, sessionKey string, action string) error {
	return nil
}
