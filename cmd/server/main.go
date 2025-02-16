package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DanikDaraboz/StoreProject/config"
	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/DanikDaraboz/StoreProject/internal/repository"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/internal/routes"
	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/DanikDaraboz/StoreProject/pkg/middleware"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	router := mux.NewRouter()

	mongoClient, err := mongo.Connect(cfg.MongoURI)
	if err != nil {
		logger.ErrorLogger.Println("Failed to connect to mongo:", err)
	}

	db := mongoClient.Database("ecommerce")
	repos := repository.NewRepositories(db)
	services := services.NewServices(repos)

	templateCache, err := handlers.NewTemplateCache()
	if err != nil {
		logger.ErrorLogger.Println("Failed to load templates:", err)
		os.Exit(1)
	}

	middlewareInstance := middleware.NewMiddleware(services)

	srv := handlers.NewServer(router, mongoClient, services, templateCache, middlewareInstance)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	routes.RegisterRoutes(srv)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: srv.Router,
	}

	// Clear expired sessions from DB
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			err := srv.Services.SessionServices.ClearExpiredSessions()
			if err != nil {
				logger.WarnLogger.Println("Failed to clear expired sessions:", err)
			} else {
				logger.InfoLogger.Println("Expired sessions cleared successfully")
			}
		}
	}()

	// Graceful shutdown
	go func() {
		logger.InfoLogger.Println("Server running on port", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorLogger.Println("Server error:", err)
		}

	}()

	// Listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.ErrorLogger.Println("Shutting down server...")

	// Extra time for ongoing requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.ErrorLogger.Println("Server forced shutdown:", err)
	}

	logger.InfoLogger.Println("Graceful shutdown of the server complete!")
}