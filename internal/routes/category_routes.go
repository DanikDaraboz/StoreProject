package routes

import "github.com/DanikDaraboz/StoreProject/internal/handlers"

func RegisterCategoryRoutes(s *handlers.Server) {
	// Public category routes
	s.Router.HandleFunc("/categories", s.GetAllCategories).Methods("GET")

	// Protected (Admin-only)
	s.Router.HandleFunc("/categories", s.CreateCategory).Methods("POST")
	s.Router.HandleFunc("/categories/{id}", s.UpdateCategory).Methods("PUT")
	s.Router.HandleFunc("/categories/{id}", s.DeleteCategory).Methods("DELETE")

}
