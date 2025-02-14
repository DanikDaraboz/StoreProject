package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.ErrorLogger.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// ?

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {

}
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
