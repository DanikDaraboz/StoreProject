package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func (s *Server) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		logger.ErrorLogger.Println("Session cookie not found:", err)
		http.Error(w, "No active session", http.StatusBadRequest)
		return
	}

	sessionKey := cookie.Value
	session, err := s.Services.SessionServices.FindSession(sessionKey)
	if err != nil {
		logger.ErrorLogger.Println("Session not found:", err)
		http.Error(w, "No active session", http.StatusInternalServerError)
		return
	}

	user, err := s.Services.UserServices.GetUser(session.UserID)
	if err != nil {
		logger.ErrorLogger.Println("User not found:", err)
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	// Fetch all products
	products, err := s.Services.ProductServices.GetAllProducts()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Fetch all categories
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Prepare the template data
	data := TemplateData{
		Title:      "Home page",
		Products:   &products,
		Categories: &categories,
		User:       user,
	}

	// Retrieve the cached template
	ts, ok := s.TemplatesCache["index.html"]
	if !ok {
		logger.ErrorLogger.Println("Template index.html not found in cache")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Printf("Failed to render template: %v", err)
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
