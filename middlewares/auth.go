package middlewares

import (
	"github.com/DanikDaraboz/StoreProject/utils"
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler, client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve session cookie
		cookie, err := r.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			log.Println("No session cookie found or empty value")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Validate session
		userID, err := utils.ValidateSession(client, cookie.Value)
		if err != nil {
			log.Println("Session validation failed:", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		log.Println("Session validated. User ID:", userID)

		// Add userID to context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
