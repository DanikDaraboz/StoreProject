package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.UserServicesInterface = (*userServices)(nil)

type userServices struct {
	userRepo   repoInterface.UserRepositoryInterface
	sessionSvc interfaces.SessionServicesInterface
}

func NewUserServices(userRepo repoInterface.UserRepositoryInterface, sessionSvc interfaces.SessionServicesInterface) interfaces.UserServicesInterface {
	return &userServices{userRepo: userRepo, sessionSvc: sessionSvc}
}

func (u *userServices) RegisterUser(user *models.User) (primitive.ObjectID, error) {
	if user.Email == "" || user.Password == "" {
		return primitive.NilObjectID, errors.New("email and password required")
	}

	existingUser, err := u.userRepo.FindUserByEmail(user.Email)
	if err == nil && existingUser.ID != primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("user already exists")
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = hashPassword

	user.ID = primitive.NewObjectID()
	userID, err := u.userRepo.InsertUser(user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userID, nil
}

func (u *userServices) LoginUser(email string, password string) (string, error) {
	// Retrieve user by email
	user, err := u.userRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check password using the utility function
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	sessionKey, err := u.sessionSvc.CreateSession(user.ID)
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}

func (u *userServices) LogoutUser(sessionKey string) error {
	if sessionKey == "" {
		return errors.New("sessionKey empty")
	}

	err := u.sessionSvc.DeleteSession(sessionKey)

	return err
}

func (u *userServices) GetUser(userID primitive.ObjectID) (*models.User, error) {
	if userID == primitive.NilObjectID {
		return &models.User{}, errors.New("userID is empty")
	}

	user, err := u.userRepo.FindUserByID(userID)

	return user, err
}

func (u *userServices) UpdateUser(user *models.User) error {
	if user.ID == primitive.NilObjectID {
		return errors.New("invalid user ID")
	}

	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	err := u.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
