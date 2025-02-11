package routes

import (
	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")   // List products
	router.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET") // Get product by ID
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST") // Add product
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")  // Update product
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE") // Delete product
}
