package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"

	"github.com/DanikDaraboz/StoreProject/database/src/models"
	"github.com/DanikDaraboz/StoreProject/database/src/utils"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerRequest RegisterRequest
	err := json.NewDecoder(r.Body).Decode(RegisterRequest{})
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "testdb"
	}
	collection := utils.DB.Database(dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newUser := models.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}
	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User Registered")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login User")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "testdb"
	}
	collection := utils.DB.Database(dbName).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, "Failed to decode user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Error with cursor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
		return
	}
}
