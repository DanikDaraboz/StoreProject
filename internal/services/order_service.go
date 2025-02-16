package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.OrderServicesInterface = (*orderServices)(nil)

type orderServices struct {
	orderRepo repoInterface.OrderRepositoryInterface
}

func NewOrderServices(orderRepo repoInterface.OrderRepositoryInterface) interfaces.OrderServicesInterface {
	return &orderServices{orderRepo: orderRepo}
}

func (o *orderServices) FindAllOrders() ([]models.Order, error) {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching orders:", err)
		return nil, err
	}

	return orders, nil
}

// TODO business logic
func (o *orderServices) GetOrderByID(id string) (*models.Order, error) {
	return o.orderRepo.FetchOrderByID(id)
}

func (o *orderServices) CreateOrder(order *models.Order) (primitive.ObjectID, error) {
	if err := validateOrder(order); err != nil {
		logger.ErrorLogger.Println("Order validation failed:", err)
		return primitive.NilObjectID, err
	}

	return o.orderRepo.InsertOrder(order)
}

func (o *orderServices) UpdateOrder(id string, order *models.Order) error {
	return o.orderRepo.UpdateOrder(id, order)
}

func (o *orderServices) DeleteOrder(id string) error {
	return o.orderRepo.RemoveOrder(id)
}

func validateOrder(order *models.Order) error {
	if order.UserID == "" {
		return errors.New("customer ID is required")
	}
	if len(order.Items) == 0 {
		return errors.New("at least one product is required")
	}
	if order.TotalPrice <= 0 {
		return errors.New("total amount must be positive")
	}
	return nil
}
