package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	"github.com/DanikDaraboz/StoreProject/database/src/models"
	"github.com/DanikDaraboz/StoreProject/database/src/utils"
)

// Claims определяет структуру данных, хранящихся в JWT
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// SecretKey - секретный ключ для подписи JWT. В production храните его в переменной окружения.
var SecretKey = []byte(os.Getenv("SECRET_KEY"))

// RegisterRequest определяет структуру запроса для регистрации
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest определяет структуру запроса для логина
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler обрабатывает запросы на регистрацию
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var registerRequest RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		log.Printf("Ошибка при разборе JSON: %v", err)
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	log.Printf("Получен запрос на регистрацию: Email=%s", registerRequest.Email) // Добавлено логирование

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Ошибка при хешировании пароля: %v", err)
		http.Error(w, "Ошибка при регистрации", http.StatusInternalServerError)
		return
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "Sport_Store" // Default db name
	}

	collection := utils.DB.Database(dbName).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser models.User
	err = collection.FindOne(ctx, bson.M{"email": registerRequest.Email}).Decode(&existingUser)
	if err == nil {
		http.Error(w, "Email уже существует", http.StatusBadRequest)
		return
	} else if err != mongo.ErrNoDocuments {
		log.Printf("Ошибка при проверке существования пользователя: %v", err)
		http.Error(w, "Ошибка при регистрации", http.StatusInternalServerError)
		return
	}

	// Create new user
	user := models.User{
		Email:    registerRequest.Email,
		Password: string(hashedPassword),
	}
	log.Printf("Создается пользователь: %+v", user) // Добавлено логирование

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Ошибка при вставке пользователя: %v", err)
		http.Error(w, "Ошибка при регистрации", http.StatusInternalServerError)
		return
	}

	log.Println("Пользователь успешно зарегистрирован") // Добавлено логирование

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Пользователь успешно зарегистрирован"})
}

// LoginHandler обрабатывает запросы на логин
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Printf("Ошибка при разборе JSON: %v", err)
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "Sport_Store" // Default db name
	}

	collection := utils.DB.Database(dbName).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		} else {
			log.Printf("Ошибка при получении пользователя: %v", err)
			http.Error(w, "Ошибка при входе", http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(loginRequest.Email)
	if err != nil {
		log.Printf("Ошибка при создании токена: %v", err)
		http.Error(w, "Ошибка при создании токена", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func generateToken(email string) (string, error) {
	claims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// GetUsersHandler обрабатывает запросы на получение списка пользователей (пример)
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Пример: просто возвращаем заглушку
	users := []struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{
		{ID: "1", Email: "test@example.com"},
		{ID: "2", Email: "test2@example.com"},
	}

	json.NewEncoder(w).Encode(users)
}

// GetCategoriesHandler обрабатывает запросы на получение списка категорий
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories := []string{"Football", "Tennis", "Basketball", "Baseball", "Running", "Swimming"}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		log.Printf("Ошибка при кодировании категорий в JSON: %v", err)
		http.Error(w, "Ошибка при получении категорий", http.StatusInternalServerError)
		return
	}
}

// GetProductsHandler обрабатывает запросы на получение списка продуктов
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "Sport_Store" // Default db name
	}

	collection := utils.DB.Database(dbName).Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pagination parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10 // Default limit
	offset := 0 // Default offset

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("Неверное значение limit: %v", err)
			http.Error(w, "Неверное значение limit", http.StatusBadRequest)
			return
		}
		if l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("Неверное значение offset: %v", err)
			http.Error(w, "Неверное значение offset", http.StatusBadRequest)
			return
		}
		if o >= 0 {
			offset = o
		}
	}

	// Count total number of products
	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Ошибка при подсчете продуктов: %v", err)
		http.Error(w, "Ошибка при получении продуктов", http.StatusInternalServerError)
		return
	}

	// Set options for Find to include limit and offset for pagination
	findOptions := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("Ошибка при получении продуктов: %v", err)
		http.Error(w, "Ошибка при получении продуктов", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err = cursor.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Ошибка при декодировании продукта: %v", err)
			http.Error(w, "Ошибка при декодировании продукта", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Ошибка курсора: %v", err)
		http.Error(w, "Ошибка курсора", http.StatusInternalServerError)
		return
	}

	// Return products along with total count for pagination on client side
	response := map[string]interface{}{
		"total":    total,
		"products": products,
		"limit":    limit,
		"offset":   offset,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Ошибка при кодировании продуктов в JSON: %v", err)
		http.Error(w, "Ошибка при кодировании продуктов в JSON", http.StatusInternalServerError)
		return
	}
}

// GetProductByIDHandler обрабатывает запросы на получение продукта по ID
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract product ID from the URL path using mux.Vars
	vars := mux.Vars(r)
	productID, ok := vars["id"]
	if !ok {
		http.Error(w, "Не указан ID продукта", http.StatusBadRequest)
		return
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "Sport_Store" // Default db name
	}

	collection := utils.DB.Database(dbName).Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert the product ID from string to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Printf("Неверный формат ID продукта: %v", err)
		http.Error(w, "Неверный формат ID продукта", http.StatusBadRequest)
		return
	}

	// Query to find the product with the matching ID
	var product models.Product
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		log.Printf("Ошибка при получении продукта: %v", err)
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}

	// Encode the product as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("Ошибка при кодировании продукта: %v", err)
		http.Error(w, "Ошибка при кодировании продукта", http.StatusInternalServerError)
		return
	}
}
