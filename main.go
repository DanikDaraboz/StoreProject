package main

import (
	"log"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/db"
	"github.com/DanikDaraboz/StoreProject/middlewares"
	"github.com/DanikDaraboz/StoreProject/routes"
)

func main() {
	uri := "mongodb+srv://nero:123@neromongo.2b3kf.mongodb.net/?retryWrites=true&w=majority&appName=NeroMongo"

	// Connect to MongoDB
	dbConn := db.Connect(uri)
	defer dbConn.Disconnect(nil)

	// Public routes
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		routes.LoginHandler(w, r, dbConn)
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		routes.LogoutHandler(w, r, dbConn)
	})

	// Protected routes
	http.Handle("/dashboard", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.DashboardHandler(w, r, dbConn)
	}), dbConn))

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
