package handlers

import (
	"fmt"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home handler reached!")
	data := templateData{
		Title: "Home page",

		Products: []models.Product{
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Laptop", Price: 999.99, Images: []string{"static/images/1.jpg"}},
			{ID: primitive.NewObjectID(), Name: "Smartphone", Price: 699.99, Images: []string{"/static/images/2.jpg"}},
		},
	}

	ts, ok := s.TemplatesCache["index.html"]
	if !ok {
		fmt.Println("Template not found!")
		http.Error(w, "The template does not exist", http.StatusInternalServerError)
		return
	}

	fmt.Println("Template found. Rendering now...")

	err := ts.Execute(w, data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	fmt.Println("Template rendered successfully!")
}
