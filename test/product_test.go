package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/internal/routes"
	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/gorilla/mux"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var srv *handlers.Server // Declare global server instance

func TestMain(m *testing.M) {
	// MongoDB connection setup for tests
	mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongoDriver.Connect(context.Background(), clientOptions)
	if err != nil {
		panic("Failed to connect to MongoDB for tests: " + err.Error())
	}

	// Initialize router
	router := mux.NewRouter()

	// Initialize test services with mock or real repositories
	testServices := &services.Services{
		ProductServices: services.NewProductServices(nil), // Replace nil with a mock repo
		OrderServices:   services.NewOrderServices(nil),   // Replace nil with a mock repo
	}

	templateCache, err := handlers.NewTemplateCache()

	// Create a Server instance for testing
	srv = handlers.NewServer(router, mongoClient, testServices, templateCache)

	// Register routes using the new Server instance
	routes.RegisterRoutes(srv)

	// Run all tests
	exitCode := m.Run()

	// Cleanup
	mongo.Client.Disconnect(context.Background())
	os.Exit(exitCode)
}

// Helper function to execute test requests
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)
	return rr
}

// Helper to check response status
func checkStatus(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected status %d but got %d", expected, actual)
	}
}

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
