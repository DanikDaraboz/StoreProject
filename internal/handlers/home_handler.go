package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func (s *Server) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	// Fetch all products
	products, err := s.Services.ProductServices.GetAllProducts()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "Home page",
		Products:   &products,
		Categories: &categories,
	}

	// Render the cached template
	ts := s.TemplatesCache["index.html"]
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.InfoLogger.Println("Home page rendered successfully.")
}

func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := s.DB.Ping(ctx, nil)
	if err != nil {
		http.Error(w, "MongoDB not connected", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MongoDB is connected!"))
}
