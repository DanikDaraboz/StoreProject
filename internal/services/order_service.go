package services

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllOrders() ([]models.Order, error) {
	orders, err := mongo.GetOrders()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching orders:", err)
		return nil, err
	}

	return orders, nil
}

// TODO business logic
func GetOrderByID(id string) (models.Order, error) {
	return mongo.FetchOrderByID(id)
}

// TODO check for nil order?
func CreateOrder(order models.Order) (primitive.ObjectID, error) {
	return mongo.InsertOrder(order)
}

func UpdateOrder(id string, order models.Order) error {
	return mongo.UpdateOrder(id, order)
}

func DeleteOrder(id string) error {
	return mongo.RemoveOrder(id)
}
