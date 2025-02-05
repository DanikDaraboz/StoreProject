package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DanikDaraboz/StoreProject/database/src/routes"
	"github.com/DanikDaraboz/StoreProject/database/src/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load .env file
	if err := godotenv.Load("/Users/bellyashsh/GolandProjects/StoreProject/database/.env"); err != nil { // Пример
		log.Println("No .env file found")
	}

	// Connect to MongoDB
	utils.ConnectDB()
	log.Println("Подключение к MongoDB успешно установлено") // Добавлено логирование

	// Create a new router
	router := mux.NewRouter()

	// Маршрут для статических файлов
	staticDir := "./static"
	log.Printf("Статическая директория: %s", staticDir)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Setup routes
	routes.SetupRoutes(router)

	// CORS configuration
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000" // Значение по умолчанию
	}
	origins := []string{allowedOrigins}

	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true, // Включите для отладки, но выключите в production
	})

	// Apply CORS middleware to the router
	handler := c.Handler(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
