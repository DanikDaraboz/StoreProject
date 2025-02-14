package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := s.DB.Ping(ctx, nil)
	if err != nil {
		http.Error(w, "MongoDB not connected", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MongoDB is connected!"))
}
