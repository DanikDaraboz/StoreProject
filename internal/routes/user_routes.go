package routes

import (
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
)

func RegisterUserRoutes(s *handlers.Server) {
	// Public user auth routes
	s.Router.HandleFunc("/login", s.RenderLoginPage).Methods("GET")       // Render login page
	s.Router.HandleFunc("/login", s.LoginUser).Methods("POST")            // Process login
	s.Router.HandleFunc("/register", s.RenderRegisterPage).Methods("GET") // Render registration form
	s.Router.HandleFunc("/register", s.RegisterUser).Methods("POST")      // Process registration
	s.Router.HandleFunc("/logout", s.Logout).Methods("GET", "POST")

	// Protected user routes (User-only)
	s.Router.Handle("/user", s.Middleware.AuthMiddleware(http.HandlerFunc(s.RenderUserProfilePage))).Methods("GET") // View user profile
	s.Router.Handle("/user", s.Middleware.AuthMiddleware(http.HandlerFunc(s.UpdateUser))).Methods("PUT")            // Update user profile

	// Protected (Admin-only)
	s.Router.Handle("/admin", s.Middleware.AuthMiddleware(s.Middleware.AdminOnlyMiddleware(http.HandlerFunc(s.RenderAdminPage)))).Methods("GET")
}
