package interfaces

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductServicesInterface interface {
	GetAllProducts() ([]map[string]interface{}, error)
	GetProductByID(id string) (models.Product, error)
	CreateProduct(product models.Product) error
	UpdateProduct(id string, product models.Product) error
	DeleteProduct(id string) error
}

type OrderServicesInterface interface {
	FindAllOrders() ([]models.Order, error)
	GetOrderByID(id string) (models.Order, error)
	CreateOrder(order models.Order) (primitive.ObjectID, error)
	UpdateOrder(id string, order models.Order) error
	DeleteOrder(id string) error
}

type CartServicesInterface interface {
	AddItemToCart(userID string, item models.CartItem) error
	GetCartItems(userID string) ([]models.CartItem, error)
	RemoveItemFromCart(userID string, itemID string) error
	ClearCart(userID string) error	
	UpdateCartItem(userID string, itemID string, quantity int) error
}

type UserServicesInterface interface {
	RegisterUser(user models.User) (primitive.ObjectID, error)
	LoginUser(email string, password string) (string, error) // Returns session key
	LogoutUser(sessionKey string) error
	GetUserByID(id string) (models.User, error)
	UpdateUser(id string, user models.User) error
	ValidateSession(sessionKey string) (models.User, error)
}