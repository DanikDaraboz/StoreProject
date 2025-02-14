package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/gorilla/mux"
)

func (s *Server) AddCartItem(w http.ResponseWriter, r *http.Request) {
	userID := "get" // TODO

	var item models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		logger.ErrorLogger.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	if err := s.Services.CartServices.AddItemToCart(userID, item); err != nil {
		logger.ErrorLogger.Println("Failed to add item to cart:", err)
		http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetCartItems(w http.ResponseWriter, r *http.Request) {
	userID := "get" // TODO

	var items []models.CartItem
	items, err := s.Services.CartServices.GetCartItems(userID)
	if err != nil {
		logger.ErrorLogger.Println("Failed to fetch cart items:", err)
		http.Error(w, "Failed to fetch cart items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (s *Server) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	userID := "get" // TODO

	vars := mux.Vars(r)
	itemID := vars["ProductID"]

	if err := s.Services.CartServices.RemoveItemFromCart(userID, itemID); err != nil {
		logger.ErrorLogger.Println("Failed to remove cart item:", err)
		http.Error(w, "Failed to remove cart item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}


// ClearCart?