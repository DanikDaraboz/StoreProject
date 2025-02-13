package services

import (
	"github.com/DanikDaraboz/StoreProject/internal/repository"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
)

type Services struct {
	ProductServices interfaces.ProductServicesInterface
	OrderServices   interfaces.OrderServicesInterface
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		ProductServices: NewProductServices(repos.ProductRepo),
		OrderServices:   NewOrderServices(repos.OrderRepo),
	}
}
