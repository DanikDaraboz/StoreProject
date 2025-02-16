package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/gorilla/mux"
)

func (s *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := s.Services.CategoryService.CreateCategory(category)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"category_id": id.Hex()})
}

func (s *Server) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (s *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var category *models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := s.Services.CategoryService.UpdateCategory(id, category); err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := s.Services.CategoryService.DeleteCategory(id); err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
