package interfaces

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductServicesInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(id string, product *models.Product) error
	DeleteProduct(id string) error
}

type OrderServicesInterface interface {
	FindAllOrders() ([]models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
	CreateOrder(order *models.Order) (primitive.ObjectID, error)
	UpdateOrder(id string, order *models.Order) error
	DeleteOrder(id string) error
}

type CartServicesInterface interface {
	GetCart(userID primitive.ObjectID) (*models.Cart, error)
	UpdateCartItem(userID primitive.ObjectID, productID primitive.ObjectID, quantity int) error
	AddItemToCart(userID primitive.ObjectID, item *models.CartItem) error
	GetCartItems(userID primitive.ObjectID) ([]models.CartItem, error)
	RemoveItemFromCart(userID primitive.ObjectID, itemID primitive.ObjectID) error
	ClearCart(userID primitive.ObjectID) error
	UpdateCart(cart *models.Cart) error
}

type UserServicesInterface interface {
	RegisterUser(user *models.User) (primitive.ObjectID, error)
	LoginUser(email string, password string) (string, error) // Returns session key
	LogoutUser(sessionKey string) error
	GetUser(userID primitive.ObjectID) (*models.User, error)
	UpdateUser(user *models.User) error
}

type SessionServicesInterface interface {
	CreateSession(userID primitive.ObjectID) (string, error)
	FindSession(sessionID string) (*models.Session, error)
	DeleteSession(sessionID string) error
	ClearExpiredSessions() error
}

type CategoryServicesInterface interface {
	CreateCategory(category *models.Category) (primitive.ObjectID, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id string) (*models.Category, error)
	UpdateCategory(id string, category *models.Category) error
	DeleteCategory(id string) error
}
