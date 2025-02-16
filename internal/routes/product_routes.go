package routes

import (
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
)

func RegisterProductRoutes(s *handlers.Server) {
	// Public product routes
	s.Router.HandleFunc("/products", s.RenderProductsPage).Methods("GET")            // Render product list
	s.Router.HandleFunc("/products/{id}", s.RenderProductDetailsPage).Methods("GET") // Render product details

	// Protected product routes (Admin-only)
	s.Router.Handle("/products", s.Middleware.AuthMiddleware(http.HandlerFunc(s.CreateProduct))).Methods("POST")        // Add product
	s.Router.Handle("/products/{id}", s.Middleware.AuthMiddleware(http.HandlerFunc(s.UpdateProduct))).Methods("PUT")    // Update product
	s.Router.Handle("/products/{id}", s.Middleware.AuthMiddleware(http.HandlerFunc(s.DeleteProduct))).Methods("DELETE") // Delete product
}
