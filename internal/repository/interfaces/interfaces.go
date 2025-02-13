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
