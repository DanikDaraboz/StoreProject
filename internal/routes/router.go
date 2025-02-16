package routes

import (
	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/DanikDaraboz/StoreProject/pkg/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes(s *handlers.Server) {

	s.Router.Use(middleware.CORSMiddleware)
	s.Router.Use(middleware.LoggerMiddleware)
	s.Router.Use(middleware.RecoveryMiddleware)

	// TODO
	// ClearCart, UpdateCartItemQuantity

	// Public Routes
	s.Router.HandleFunc("/", s.RenderHomePage).Methods("GET")
	s.Router.HandleFunc("/health", s.HealthCheckHandler).Methods("GET")

	RegisterProductRoutes(s)
	RegisterOrderRoutes(s)
	RegisterCartRoutes(s)
	RegisterUserRoutes(s)
	RegisterCategoryRoutes(s)

	logger.InfoLogger.Println("Registered routes:")
	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			logger.InfoLogger.Println("Route:", path)
		}
		return nil
	})
}
