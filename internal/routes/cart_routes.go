package routes

import "github.com/DanikDaraboz/StoreProject/internal/handlers"

func RegisterCartRoutes(s *handlers.Server) {
	// Public cart routes
	s.Router.HandleFunc("/cart", s.AddCartItem).Methods("POST")
	s.Router.HandleFunc("/cart", s.RenderCartPage).Methods("GET")     
	s.Router.HandleFunc("/cart/clear", s.ClearCart).Methods("DELETE")
	s.Router.HandleFunc("/cart/{id}", s.DeleteCartItem).Methods("DELETE")
	s.Router.HandleFunc("/cart/item", s.UpdateCartItem).Methods("PUT")
}
