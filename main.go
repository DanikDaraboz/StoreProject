package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/DanikDaraboz/StoreProject/database/src/handlers"
	"github.com/DanikDaraboz/StoreProject/database/src/middleware"
	"github.com/DanikDaraboz/StoreProject/database/src/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

var templates = template.Must(template.ParseFiles(
	"templates/layout.html",
	"templates/partials/header.html",
	"templates/partials/footer.html",
	"templates/products.html",
	"templates/product_details.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template error:", err)
	}
}

func main() {
	err := godotenv.Load("database/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.ConnectDB()
	log.Println("Application started")

	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")
	log.Printf("DB_NAME: %s", dbName)
	log.Printf("SECRET_KEY: %s", secretKey)
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "layout.html", map[string]interface{}{
			"Title": "Home",
		})
	})

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "layout.html", map[string]interface{}{
			"Title": "Products",
		})
	})

	r.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "layout.html", map[string]interface{}{
			"Title": "Product Details",
		})
	})

	r.HandleFunc("/api/products", handlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", handlers.GetProductByIDHandler).Methods("GET")
	r.Handle("/api/users", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetUsersHandler))).Methods("GET")
	r.HandleFunc("/api/users/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/users/login", handlers.LoginHandler).Methods("POST")
	loggedRouter := middlewares.LoggingMiddleware(r)
	http.Handle("/", loggedRouter)

	pwd, _ := os.Getwd()
	log.Println("Текущая рабочая директория:", pwd)
	log.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
