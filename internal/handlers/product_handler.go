package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/gorilla/mux"
)

func (s *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.Services.ProductServices.GetAllProducts()
	if err != nil {
		logger.ErrorLogger.Println("Error fetching products:", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(products)
}

func (s *Server) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	product, err := s.Services.ProductServices.GetProductByID(productID)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)
}

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.Services.ProductServices.CreateProduct(product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.Services.ProductServices.UpdateProduct(productID, updatedProduct); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedProduct)
}

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	if err := s.Services.ProductServices.DeleteProduct(productID); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) RenderProductsPage(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) RenderProductDetailsPage(w http.ResponseWriter, r *http.Request) {

}
