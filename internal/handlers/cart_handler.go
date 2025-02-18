package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) AddCartItem(w http.ResponseWriter, r *http.Request) {
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
	logger.ErrorLogger.Println("User found:", user)

	var item *models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		logger.ErrorLogger.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	if err := s.Services.CartServices.AddItemToCart(user.ID, item); err != nil {
		logger.ErrorLogger.Println("Failed to add item to cart:", err)
		http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item added to cart"})
}

func (s *Server) GetCartItems(w http.ResponseWriter, r *http.Request) {
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

	var items []models.CartItem
	items, err = s.Services.CartServices.GetCartItems(user.ID)
	if err != nil {
		logger.ErrorLogger.Println("Failed to fetch cart items:", err)
		http.Error(w, "Failed to fetch cart items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (s *Server) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.ErrorLogger.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Convert the product ID hex string to a primitive.ObjectID
	prodID, err := primitive.ObjectIDFromHex(input.ProductID)
	if err != nil {
		logger.ErrorLogger.Println("Invalid product ID:", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := s.Services.CartServices.UpdateCartItem(user.ID, prodID, input.Quantity); err != nil {
		logger.ErrorLogger.Println("Failed to update cart item:", err)
		http.Error(w, "Failed to update cart item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cart item updated successfully"})
}

func (s *Server) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	itemID := vars["id"]
	logger.InfoLogger.Println("itemID:", itemID)

	if err = s.Services.CartServices.RemoveItemFromCart(user.ID, itemID); err != nil {
		logger.ErrorLogger.Println("Failed to remove cart item:", err)
		http.Error(w, "Failed to remove cart item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) ClearCart(w http.ResponseWriter, r *http.Request) {
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

	// Clear the user's cart
	if err := s.Services.CartServices.ClearCart(user.ID); err != nil {
		logger.ErrorLogger.Println("Failed to clear cart:", err)
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cart cleared successfully"})
}

func (s *Server) RenderCartPage(w http.ResponseWriter, r *http.Request) {
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

	// Retrieve the cart using the GetCart service
	cart, err := s.Services.CartServices.GetCart(user.ID)
	if err != nil {
		logger.ErrorLogger.Println("Error retrieving cart:", err)
		// If no cart exists, you may want to create an empty cart for display
		cart = &models.Cart{
			UserID:     user.ID,
			Items:      []models.CartItem{},
			TotalPrice: 0.0,
		}
	}

	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "Your Cart",
		Cart:       cart,
		User:       user,
		Categories: &categories,
	}

	ts, ok := s.TemplatesCache["cart.html"]
	if !ok {
		logger.ErrorLogger.Println("Template cart.html not found in cache")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Error rendering cart template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
