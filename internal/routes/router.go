package routes

import (
	"fmt"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(s *handlers.Server) {

	// TODO middleware

	s.Router.HandleFunc("/", s.Home).Methods("GET")

	s.Router.HandleFunc("/products", s.GetProducts).Methods("GET")           // List products
	s.Router.HandleFunc("/products/{id}", s.GetProductByID).Methods("GET")   // Get product by ID
	s.Router.HandleFunc("/products", s.CreateProduct).Methods("POST")        // Add product
	s.Router.HandleFunc("/products/{id}", s.UpdateProduct).Methods("PUT")    // Update product
	s.Router.HandleFunc("/products/{id}", s.DeleteProduct).Methods("DELETE") // Delete product

	s.Router.HandleFunc("/orders", s.GetOrdersHandler).Methods("GET")           // List all orders
	s.Router.HandleFunc("/orders/{id}", s.GetOrderByIDHandler).Methods("GET")   // Get order by ID
	s.Router.HandleFunc("/orders", s.CreateOrderHandler).Methods("POST")        // Create new order
	s.Router.HandleFunc("/orders/{id}", s.UpdateOrderHandler).Methods("PUT")    // Update order
	s.Router.HandleFunc("/orders/{id}", s.DeleteOrderHandler).Methods("DELETE") // Delete order

	s.Router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	fmt.Println("Registered routes:")
	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("Route:", path)
		}
		return nil
	})
}
