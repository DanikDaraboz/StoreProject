package routes

import (
	"github.com/DanikDaraboz/StoreProject/database/src/handlers"
	"github.com/DanikDaraboz/StoreProject/database/src/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/users/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/users/login", handlers.LoginHandler).Methods("POST")

	router.Use(middlewares.LoggingMiddleware)

	router.HandleFunc("/api/products", handlers.GetProductsHandler).Methods("GET")
	router.HandleFunc("/api/products/{id}", handlers.GetProductByIDHandler).Methods("GET")
	router.HandleFunc("/api/categories", handlers.GetCategoriesHandler).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	api.HandleFunc("/users/me", handlers.GetMeHandler).Methods("GET")

	admin := api.PathPrefix("").Subrouter()
	admin.Use(middlewares.AdminMiddleware)
	admin.HandleFunc("/products", handlers.CreateProductHandler).Methods("POST")
}
