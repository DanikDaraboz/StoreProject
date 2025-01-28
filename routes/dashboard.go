package routes

import (
	"github.com/DanikDaraboz/StoreProject/middlewares"

	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	// Extract userID from context
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	log.Println("userID:", userID)
	if !ok {
		log.Println("User not found in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, "templates/login.html")
}
