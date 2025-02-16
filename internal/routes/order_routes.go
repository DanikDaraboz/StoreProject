package routes

import (
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
)

func RegisterOrderRoutes(s *handlers.Server) {
	// Protected order routes (User-only)
	s.Router.Handle("/orders", s.Middleware.AuthMiddleware(http.HandlerFunc(s.CreateOrderHandler))).Methods("POST")        // Create a new order
	s.Router.Handle("/orders", s.Middleware.AuthMiddleware(http.HandlerFunc(s.GetOrdersHandler))).Methods("GET")           // Get all orders of user
	s.Router.Handle("/orders/{id}", s.Middleware.AuthMiddleware(http.HandlerFunc(s.GetOrderByIDHandler))).Methods("GET")   // Get details of the order
	s.Router.Handle("/orders/{id}", s.Middleware.AuthMiddleware(http.HandlerFunc(s.UpdateOrderHandler))).Methods("PUT")    // Update an order (status?)
	s.Router.Handle("/orders/{id}", s.Middleware.AuthMiddleware(http.HandlerFunc(s.DeleteOrderHandler))).Methods("DELETE") // Cancel an order

}
