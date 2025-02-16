package routes

import "github.com/DanikDaraboz/StoreProject/internal/handlers"

func RegisterCartRoutes(s *handlers.Server) {
	// Public cart routes
	s.Router.HandleFunc("/cart", s.AddCartItem).Methods("POST")           // Add item to the cart
	s.Router.HandleFunc("/cart", s.RenderCartPage).Methods("GET")         // Render the cart page
	s.Router.HandleFunc("/cart/{id}", s.DeleteCartItem).Methods("DELETE") // Remove item from the cart
}
