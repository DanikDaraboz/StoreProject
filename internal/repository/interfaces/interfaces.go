package interfaces

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepositoryInterface interface {
	GetProducts() ([]map[string]interface{}, error)
	FetchProductByID(id string) (models.Product, error)
	InsertProduct(product models.Product) error
	UpdateProduct(id string, product models.Product) error
	RemoveProduct(id string) error
}

type OrderRepositoryInterface interface {
	GetOrders() ([]models.Order, error)
	FetchOrderByID(id string) (models.Order, error)
	InsertOrder(order models.Order) (primitive.ObjectID, error)
	UpdateOrder(id string, order models.Order) error
	RemoveOrder(id string) error
}

type CartRepositoryInterface interface {
	InsertCart(cart models.Cart) error
	InsertCartItem(cart models.Cart, item models.CartItem) error
	FindCartByUserID(userID primitive.ObjectID) (models.Cart, error)
	FindCartItems(userID primitive.ObjectID) ([]models.CartItem, error)
	UpdateCart(cart models.Cart) error
	UpdateCartItemQuantity(userID primitive.ObjectID, itemID primitive.ObjectID, quantity int) error
	DeleteCartItem(userID primitive.ObjectID, itemID primitive.ObjectID) error
	ClearCart(userID primitive.ObjectID) error
}

type UserRepositoryInterface interface {
	InsertUser(user models.User) (primitive.ObjectID, error)
	FindUserByID(id primitive.ObjectID) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	UpdateUser(id primitive.ObjectID, user models.User) error
	InsertSession(userID primitive.ObjectID, sessionKey string) error
	FindUserBySessionKey(sessionKey string) (models.User, error)
	DeleteSession(sessionKey string) error
}
