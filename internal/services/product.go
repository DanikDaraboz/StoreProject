package services

import (
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func GetAllProducts() ([]map[string]interface{}, error) {
	products, err := mongo.GetProducts()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching products:", err)
		return nil, err
	}
	return products, nil
}
