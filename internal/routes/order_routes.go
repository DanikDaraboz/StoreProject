package routes

import (
	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/gorilla/mux"
)

// RegisterOrderRoutes registers all order-related endpoints
func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/orders", handlers.GetOrdersHandler).Methods("GET") // List all orders
	router.HandleFunc("/orders/{id}", handlers.GetOrderByIDHandler).Methods("GET")  // Get order by ID
	router.HandleFunc("/orders", handlers.CreateOrderHandler).Methods("POST")       // Create new order
	router.HandleFunc("/orders/{id}", handlers.UpdateOrderHandler).Methods("PUT")   // Update order
	router.HandleFunc("/orders/{id}", handlers.DeleteOrderHandler).Methods("DELETE") // Delete order
}
