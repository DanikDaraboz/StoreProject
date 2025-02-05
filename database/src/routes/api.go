package routes

import (
	"github.com/DanikDaraboz/StoreProject/database/src/handlers"
	"github.com/DanikDaraboz/StoreProject/database/src/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	// Public routes
	router.HandleFunc("/api/users/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/users/login", handlers.LoginHandler).Methods("POST")

	// Middleware для логирования (apply to all routes)
	router.Use(middlewares.LoggingMiddleware)

	// Protected routes - require authentication
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware) // Apply authentication middlewares

	api.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	api.HandleFunc("/products", handlers.GetProductsHandler).Methods("GET")
	api.HandleFunc("/products/{id}", handlers.GetProductByIDHandler).Methods("GET")
	api.HandleFunc("/categories", handlers.GetCategoriesHandler).Methods("GET")

	// Дополнительные маршруты (см. ниже)
}
