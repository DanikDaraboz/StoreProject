package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func GetAllProducts() ([]map[string]interface{}, error) {
	products, err := mongo.GetProducts()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching products:", err)
		return nil, err
	}

	// TODO filtering

	return products, nil
}

func GetProductByID(id string) (models.Product, error) {
	return mongo.FetchProductByID(id)
}

func CreateProduct(product models.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("invalid product data")
	}
	return mongo.InsertProduct(product)
}

func UpdateProduct(id string, product models.Product) error {
	return mongo.UpdateProduct(id, product)
}

func DeleteProduct(id string) error {
	return mongo.RemoveProduct(id)
}
