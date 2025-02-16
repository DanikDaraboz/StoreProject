package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/gorilla/mux"
)

func (s *Server) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := s.Services.OrderServices.FindAllOrders()
	if err != nil {
		logger.ErrorLogger.Println("Failed to fetch orders:", err)
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (s *Server) GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	order, err := s.Services.OrderServices.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (s *Server) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		logger.ErrorLogger.Println("Failed to decode order:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	insertedID, err := s.Services.OrderServices.CreateOrder(order)
	if err != nil {
		logger.ErrorLogger.Println("Failed to create order:", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	order.ID = insertedID // Assign returned ObjectID

	logger.InfoLogger.Println("Created orderID:", order.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(order)
}

func (s *Server) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	var updatedOrder *models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.Services.OrderServices.UpdateOrder(orderID, updatedOrder); err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedOrder)
}

func (s *Server) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	if err := s.Services.OrderServices.DeleteOrder(orderID); err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
