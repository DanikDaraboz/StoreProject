package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.Services.ProductServices.GetProducts("")
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
	var product *models.Product
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

	var updatedProduct struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		Category    string  `json:"category"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := models.Product{
		ID:          objectID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		Stock:       updatedProduct.Stock,
		Category:    updatedProduct.Category,
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}

	if err := s.Services.ProductServices.UpdateProduct(productID, &product); err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
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
	categoryID := r.URL.Query().Get("category_id")

	// Retrieve products from the service.
	products, err := s.Services.ProductServices.GetProducts(categoryID)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	// Retrieve categories for navigation.
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:      "All Products",
		Products:   &products,
		Categories: &categories,
	}

	ts, ok := s.TemplatesCache["allproducts.html"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	if err := ts.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (s *Server) RenderProductDetailsPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	var user *models.User = nil

	if err != nil {
		// Session cookie not found
		logger.ErrorLogger.Println("Session cookie not found:", err)
	} else {
		sessionKey := cookie.Value
		session, err := s.Services.SessionServices.FindSession(sessionKey)
		if err != nil {
			logger.ErrorLogger.Println("Session not found:", err)
		} else {
			// Retrieve the user associated with this session
			user, err = s.Services.UserServices.GetUser(session.UserID)
			if err != nil {
				logger.ErrorLogger.Println("User not found:", err)
			}
		}
	}

	// Fetch all categories for the navigation menu
	categories, err := s.Services.CategoryService.GetAllCategories()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to fetch categories: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	productID := vars["id"]

	// Retrieve product details using the product ID
	product, err := s.Services.ProductServices.GetProductByID(productID)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	data := TemplateData{
		Title:      "Product Page",
		User:       user,
		Product:    product,
		Categories: &categories,
	}

	ts := s.TemplatesCache["product.html"]
	if err := ts.Execute(w, data); err != nil {
		logger.ErrorLogger.Println("Failed to render template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
