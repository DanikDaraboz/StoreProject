package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/DanikDaraboz/StoreProject/database/src/models"
	"github.com/DanikDaraboz/StoreProject/database/src/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value("email").(string)
		if !ok {
			http.Error(w, "Не удалось получить email пользователя из токена", http.StatusInternalServerError)
			return
		}

		user, err := getUserByEmail(email)
		if err != nil {
			http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
			return
		}

		if user.Role != "admin" {
			http.Error(w, "Недостаточно прав", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func getUserByEmail(email string) (models.User, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "Sport_Store"
	}

	collection := utils.DB.Database(dbName).Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
