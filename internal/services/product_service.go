package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.ProductServicesInterface = (*productService)(nil)

type productService struct {
	productRepo repoInterface.ProductRepositoryInterface
}

func NewProductServices(productRepo repoInterface.ProductRepositoryInterface) interfaces.ProductServicesInterface {
	return &productService{productRepo: productRepo}
}

func (p *productService) GetProducts(categoryID string) ([]models.Product, error) {
	// Call the repository to get products with optional category filter
	products, err := p.productRepo.GetProducts(categoryID)
	if err != nil {
		logger.ErrorLogger.Println("Error fetching products:", err)
		return nil, err
	}

	// TODO:
	// - Validate product categories against predefined categories.
	// - Ensure product.Images contains valid URLs.
	// - Implement pagination for large datasets.
	// - Add safe stock updates to prevent race conditions.

	return products, nil
}
func (p *productService) GetProductByID(id string) (*models.Product, error) {
	if id == "" {
		return &models.Product{}, errors.New("product ID cannot be empty")
	}

	return p.productRepo.FetchProductByID(id)
}

func (p *productService) CreateProduct(product *models.Product) error {
	if err := validateProduct(product); err != nil {
		logger.ErrorLogger.Println("Product validation failed:", err)
		return err
	}

	return p.productRepo.InsertProduct(product)
}

func (p *productService) UpdateProduct(id string, product *models.Product) error {
	if id == "" {
		return errors.New("product ID cannot be empty")
	}

	if err := validateProduct(product); err != nil {
		return err
	}

	// check if exist before updating
	existingProduct, err := p.productRepo.FetchProductByID(id)
	if err != nil {
		return err
	}

	if !isProductChanged(existingProduct, product) {
		logger.InfoLogger.Printf("No changes detected for product ID: %s", id)
		return nil
	}

	return p.productRepo.UpdateProduct(id, product)
}

func (p *productService) DeleteProduct(id string) error {
	if id == "" {
		return errors.New("product ID cannot be empty")
	}

	_, err := p.productRepo.FetchProductByID(id)
	if err == mongo.ErrNoDocuments {
		return errors.New("product not found, cannot delete")
	}
	if err != nil {
		return err
	}

	return p.productRepo.RemoveProduct(id)
}

func validateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	if len(product.Category) == 0 {
		return errors.New("product category is required")
	}
	return nil
}

func isProductChanged(old, new *models.Product) bool {
	if old.Name != new.Name ||
		old.Price != new.Price ||
		old.Stock != new.Stock ||
		old.Category != new.Category {
		return true
	}

	if len(old.Images) != len(new.Images) {
		return true
	}
	for i := range old.Images {
		if old.Images[i] != new.Images[i] {
			return true
		}
	}

	return false
}
