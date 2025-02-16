package services

import (
	"github.com/DanikDaraboz/StoreProject/internal/repository"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
)

type Services struct {
	ProductServices interfaces.ProductServicesInterface
	OrderServices   interfaces.OrderServicesInterface
	CartServices    interfaces.CartServicesInterface
	UserServices    interfaces.UserServicesInterface
	SessionServices interfaces.SessionServicesInterface
	CategoryService interfaces.CategoryServicesInterface
}

func NewServices(repos *repository.Repositories) *Services {
	// Initialize the session service first for UserServices
	sessionServices := NewSessionServices(repos.SessionRepo)

	return &Services{
		ProductServices: NewProductServices(repos.ProductRepo),
		OrderServices:   NewOrderServices(repos.OrderRepo),
		CartServices:    NewCartServices(repos.CartRepo),
		UserServices:    NewUserServices(repos.UserRepo, sessionServices),
		SessionServices: sessionServices, 
		CategoryService: NewCategoryService(repos.CategoryRepo),
	}
}
