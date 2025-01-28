package routes

import (
	"github.com/DanikDaraboz/StoreProject/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO
func LoginHandler(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	if r.Method == http.MethodPost {
		userID := "12345"

		// Create session
		sessionID, err := utils.CreateSession(client, userID)
		if err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

		// Set session cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "templates/login.html")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	// Retrieve session cookie
	cookie, err := r.Cookie("session_id")
	if err == nil {
		// Delete session from database
		utils.DeleteSession(client, cookie.Value)
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Expire immediately
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
