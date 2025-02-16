package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.ErrorLogger.Println("Bad request:", err)
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	sessionKey, err := s.Services.UserServices.LoginUser(loginRequest.Email, loginRequest.Password)
	if sessionKey == "" || err != nil {
		logger.WarnLogger.Println("Login failed:", err)
		http.Error(w, "Login failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "Login page",
	}

	ts := s.TemplatesCache["login.html"]
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	logger.InfoLogger.Println("Login page rendered successfully.")
}

func (s *Server) RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) RenderRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) RenderUserProfilePage(w http.ResponseWriter, r *http.Request) {

}
