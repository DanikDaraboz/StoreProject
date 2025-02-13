package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

var _ interfaces.ProductServicesInterface = (*productService)(nil)

type productService struct {
	productRepo repoInterface.ProductRepositoryInterface
}

func NewProductServices(productRepo repoInterface.ProductRepositoryInterface) interfaces.ProductServicesInterface {
	return &productService{productRepo: productRepo}
}

func (p productService) GetAllProducts() ([]map[string]interface{}, error) {
	products, err := p.productRepo.GetProducts()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching products:", err)
		return nil, err
	}

	// TODO filtering

	return products, nil
}

func (p productService) GetProductByID(id string) (models.Product, error) {
	return p.productRepo.FetchProductByID(id)
}

func (p productService) CreateProduct(product models.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("invalid product data")
	}
	return p.productRepo.InsertProduct(product)
}

func (p productService) UpdateProduct(id string, product models.Product) error {
	return p.productRepo.UpdateProduct(id, product)
}

func (p productService) DeleteProduct(id string) error {
	return p.productRepo.RemoveProduct(id)
}
