package interfaces

import (
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepositoryInterface interface {
	GetProducts(categoryID string) ([]models.Product, error)
	FetchProductByID(id string) (*models.Product, error)
	InsertProduct(product *models.Product) error
	UpdateProduct(id string, product *models.Product) error
	RemoveProduct(id string) error
}

type OrderRepositoryInterface interface {
	GetOrders() ([]models.Order, error)
	FetchOrderByID(id string) (*models.Order, error)
	InsertOrder(order *models.Order) (primitive.ObjectID, error)
	UpdateOrder(id string, order *models.Order) error
	RemoveOrder(id string) error
}

type CartRepositoryInterface interface {
	InsertCart(cart *models.Cart) error
	FindCartByUserID(userID primitive.ObjectID) (*models.Cart, error)
	UpdateCart(cart *models.Cart) error
}

type UserRepositoryInterface interface {
	InsertUser(user *models.User) (primitive.ObjectID, error)
	FindUserByID(id primitive.ObjectID) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
}

type SessionRepositoryInterface interface {
	InsertSession(sessionID string, userID primitive.ObjectID, expiresAt time.Time) error
	FindSessionByID(sessionID string) (*models.Session, error)
	DeleteSessionByID(sessionID string) error
	DeleteExpiredSessions() error
}

type CategoryRepositoryInterface interface {
	CreateCategory(category *models.Category) (primitive.ObjectID, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id string) (*models.Category, error)
	UpdateCategory(id string, category *models.Category) error
	DeleteCategory(id string) error
}
