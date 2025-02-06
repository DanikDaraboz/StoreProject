package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DanikDaraboz/StoreProject/config"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/internal/routes"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	mongo.Connect(cfg.MongoURI)

	// Init router
	mux := routes.NewRouter()

	// Start the server
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux, // TODO mux
	}

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
