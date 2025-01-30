package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DanikDaraboz/StoreProject/database/src/models"
	"github.com/DanikDaraboz/StoreProject/database/src/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func getDBCollection(collectionName string) (*mongo.Collection, context.Context, func()) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "testdb"
	}

	collection := utils.DB.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return collection, ctx, cancel
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func generateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SecretKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("Start Registration Handler")
	var registerRequest RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		log.Printf("Error on decode: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Request successfully decoded: %+v", registerRequest)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:       primitive.NewObjectID(),
		Username: registerRequest.Username,
		Password: string(hashedPassword),
	}

	collection, ctx, cancel := getDBCollection("users")
	defer cancel()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert user: %v\n", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User Registered")
	log.Println("End Registration Handler")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Printf("Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	collection, ctx, cancel := getDBCollection("users")
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"username": loginRequest.Username}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			http.Error(w, "User not found", http.StatusUnauthorized)
		} else {
			log.Printf("Failed to fetch user: %v\n", err)
			http.Error(w, "Failed to login", http.StatusInternalServerError)
		}

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := generateToken(loginRequest.Username)
	if err != nil {
		log.Printf("Failed to generate token: %v\n", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged in")

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	collection, ctx, cancel := getDBCollection("users")
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to fetch users: %v\n", err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Failed to decode user: %v\n", err)
			http.Error(w, "Failed to decode user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error with cursor: %v\n", err)
		http.Error(w, "Error with cursor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Failed to encode users to JSON: %v\n", err)
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
		return
	}
}
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "testdb"
	}

	collection := utils.DB.Database(dbName).Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to fetch products: %v\n", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Failed to decode product: %v\n", err)
			http.Error(w, "Failed to decode product", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error with cursor: %v\n", err)
		http.Error(w, "Error with cursor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("Failed to encode products to JSON: %v\n", err)
		http.Error(w, "Failed to encode products to JSON", http.StatusInternalServerError)
		return
	}
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "testdb"
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	collection := utils.DB.Database(dbName).Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to fetch product: %v\n", err)
			http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("Failed to encode product to JSON: %v\n", err)
		http.Error(w, "Failed to encode product to JSON", http.StatusInternalServerError)
		return
	}
}
