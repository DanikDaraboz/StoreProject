package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/DanikDaraboz/StoreProject/pkg/middleware"
)

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the JSON request body.
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.ErrorLogger.Println("Failed to decode JSON:", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Use the decoded values.
	sessionKey, err := s.Services.UserServices.LoginUser(input.Email, input.Password)
	if sessionKey == "" || err != nil {
		logger.WarnLogger.Println("Login failed:", err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Set the session cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionKey,
		HttpOnly: true,
		Path:     "/",
	})

	// Return the session key as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessionKey)
}

func (s *Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.ErrorLogger.Println("Failed to decode JSON:", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Create a new user with the provided data
	user := &models.User{
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	// Register the user
	userID, err := s.Services.UserServices.RegisterUser(user)
	if err != nil {
		logger.ErrorLogger.Println("Failed to register:", err)
		http.Error(w, "Failed to register", http.StatusUnauthorized)
		return
	}

	// Return the newly created user ID as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userID)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		logger.ErrorLogger.Println("Session cookie not found:", err)
		http.Error(w, "No active session", http.StatusBadRequest)
		return
	}

	sessionKey := cookie.Value

	// Call the logout service with the session key
	err = s.Services.UserServices.LogoutUser(sessionKey)
	if err != nil {
		logger.ErrorLogger.Println("Logout failed:", err)
		http.Error(w, "Logout failed", http.StatusInternalServerError)
		return
	}

	// Invalidate the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	// Return a JSON response indicating success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser *models.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		logger.ErrorLogger.Println("Invalid JSON request", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, ok := r.Context().Value(middleware.UserKey).(*models.User)
	if !ok || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	updateUser.ID = user.ID

	err := s.Services.UserServices.UpdateUser(updateUser)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updateUser)
}

func (s *Server) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "Login page",
		Categories: &categories,
	}

	ts := s.TemplatesCache["login.html"]
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	logger.InfoLogger.Println("Login page rendered successfully.")
}

func (s *Server) RenderRegisterPage(w http.ResponseWriter, r *http.Request) {
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "Register page",
		Categories: &categories,
	}

	ts := s.TemplatesCache["register.html"]
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	logger.InfoLogger.Println("Register page rendered successfully.")
}

func (s *Server) RenderUserProfilePage(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middleware.UserKey).(*models.User)
	if !ok || user == nil {
		logger.WarnLogger.Println("Failed to retrieve user from context", user)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Fetch the latest user details from the database
	updatedUser, err := s.Services.UserServices.GetUser(user.ID)
	if err != nil {
		logger.ErrorLogger.Println("Failed to fetch user details:", err)
		http.Error(w, "Failed to retrieve user profile", http.StatusInternalServerError)
		return
	}

	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "User Profile",
		Categories: &categories,
		User:       updatedUser,
	}

	tmpl, exists := s.TemplatesCache["profile.html"]
	if !exists {
		logger.ErrorLogger.Println("Template 'profile.html' not found in cache")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render user profile page:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Server) RenderAdminPage(w http.ResponseWriter, r *http.Request) {
	adminUser, ok := r.Context().Value(middleware.UserKey).(*models.User)
	if !ok || adminUser == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve all products via the product service.
	products, err := s.Services.ProductServices.GetAllProducts()
	if err != nil {
		logger.ErrorLogger.Println("Error retrieving products:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Retrieve all categories via the category service.
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Println("Error retrieving categories:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "Admin Panel",
		User:       adminUser,
		Products:   &products,
		Categories: &categories,
	}

	ts := s.TemplatesCache["admin.html"]
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
