package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/DanikDaraboz/StoreProject/internal/models"
)

func TestProductRoutes(t *testing.T) {
	t.Run("GET /products - Should return product list", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products", nil)
		response := executeRequest(req)

		checkStatus(t, http.StatusOK, response.Code)
	})

	t.Run("POST /products - Should create a product", func(t *testing.T) {
		product := models.Product{
			Name:        "Test Product",
			Description: "A sample product",
			Price:       19.99,
			Stock:       10,
			Category:    "Electronics",
		}
		payload, _ := json.Marshal(product)

		req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		response := executeRequest(req)

		checkStatus(t, http.StatusCreated, response.Code)
	})

	t.Run("GET /products/{id} - Should return a product", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products/1234567890abcdef12345678", nil)
		response := executeRequest(req)

		checkStatus(t, http.StatusNotFound, response.Code)
	})

	t.Run("PUT /products/{id} - Should update a product", func(t *testing.T) {
		update := map[string]interface{}{
			"price": 24.99,
		}
		payload, _ := json.Marshal(update)

		req, _ := http.NewRequest("PUT", "/products/1234567890abcdef12345678", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		response := executeRequest(req)

		checkStatus(t, http.StatusOK, response.Code)
	})

	t.Run("DELETE /products/{id} - Should delete a product", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/products/1234567890abcdef12345678", nil)
		response := executeRequest(req)

		checkStatus(t, http.StatusNoContent, response.Code)
	})
}
