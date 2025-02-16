package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
)

type contextKey string

const userKey contextKey = "user"
const adminKey contextKey = "admin"

var _ MiddlewareInterface = (*Middleware)(nil)

type MiddlewareInterface interface {
	AuthMiddleware(next http.Handler) http.Handler
}

type Middleware struct {
	services *services.Services
}

func NewMiddleware(srv *services.Services) MiddlewareInterface {
	return &Middleware{services: srv}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			if err == http.ErrNoCookie {
				logger.WarnLogger.Println("No sessionID in cookie")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			logger.WarnLogger.Println("Cookie read error")
			http.Error(w, "Unauthorized, couldnt read a cookies", http.StatusUnauthorized)
			return
		}

		session, err := m.services.SessionServices.FindSession(sessionID.Value)
		if err != nil {
			if errors.Is(err, errors.New("session expired")) {
				logger.WarnLogger.Println("Session retriexpired")
				http.Error(w, "Unauthorized, session expired", http.StatusUnauthorized)
				return
			}
			logger.WarnLogger.Println("Session retrieval error:", err)
			http.Error(w, "Unauthorized, couldnt fetch session", http.StatusUnauthorized)
			return
		}
		var user models.User
		user, err = m.services.UserServices.GetUser(session.UserID)
		if err != nil {
			logger.WarnLogger.Println("Couldnt get the user", err)
			http.Error(w, "Unauthorized, couldnt retrive user data", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, user)

		logger.InfoLogger.Println("User is authenticated")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WarnLogger.Printf("Request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				logger.WarnLogger.Printf("Panic recovered: %v", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// TODO Rate limit?
