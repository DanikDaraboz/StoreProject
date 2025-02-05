package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Fetch products from service
	products, err := services.GetAllProducts()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching products:", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	// Convert products to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	err := mongo.PingMongoDB()
	if err != nil {
		http.Error(w, "MongoDB not connected", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MongoDB is connected!"))
}
