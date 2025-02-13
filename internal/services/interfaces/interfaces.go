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
