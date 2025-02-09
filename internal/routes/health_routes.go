package routes

import (
	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterHealthRoutes(router *mux.Router) {
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
}
