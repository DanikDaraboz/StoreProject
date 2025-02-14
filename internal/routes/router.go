package routes

import (
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/DanikDaraboz/StoreProject/pkg/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes(s *handlers.Server) {
	s.Router.Use(middleware.CORSMiddleware)
	s.Router.Use(middleware.LoggerMiddleware)
	s.Router.Use(middleware.RecoveryMiddleware)

	// TODO
	// ClearCart, UpdateCartItemQuantity

	// Public Routes
	s.Router.HandleFunc("/", s.Home).Methods("GET")
	s.Router.HandleFunc("/health", s.HealthCheckHandler).Methods("GET")
	s.Router.HandleFunc("/products", s.GetProducts).Methods("GET")
	s.Router.HandleFunc("/products/{id}", s.GetProductByID).Methods("GET")
	s.Router.HandleFunc("/cart", s.AddCartItem).Methods("POST")           // Add item to the cart
	s.Router.HandleFunc("/cart", s.GetCartItems).Methods("GET")           // View cart items
	s.Router.HandleFunc("/cart/{id}", s.DeleteCartItem).Methods("DELETE") // Remove item from the cart
	s.Router.HandleFunc("/login", s.LoginUser).Methods("POST")            // User login
	s.Router.HandleFunc("/register", s.RegisterUser).Methods("POST")      // User registration
	s.Router.HandleFunc("/logout", s.Logout).Methods("POST")              // User logout

	// Protected product routes (Admin-only)
	s.Router.Handle("/products", middleware.AuthMiddleware(http.HandlerFunc(s.CreateProduct))).Methods("POST")        // Add a new product
	s.Router.Handle("/products/{id}", middleware.AuthMiddleware(http.HandlerFunc(s.UpdateProduct))).Methods("PUT")    // Update a product
	s.Router.Handle("/products/{id}", middleware.AuthMiddleware(http.HandlerFunc(s.DeleteProduct))).Methods("DELETE") // Delete a product

	// Protected order routes (User-only)
	s.Router.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(s.CreateOrderHandler))).Methods("POST")        // Create a new order
	s.Router.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(s.GetOrdersHandler))).Methods("GET")           // Get all orders of user
	s.Router.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(s.GetOrderByIDHandler))).Methods("GET")   // Get details of the order
	s.Router.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(s.UpdateOrderHandler))).Methods("PUT")    // Update an order (status?)
	s.Router.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(s.DeleteOrderHandler))).Methods("DELETE") // Cancel an order
	s.Router.Handle("/user", middleware.AuthMiddleware(http.HandlerFunc(s.GetUser))).Methods("GET")                      // View user profile
	s.Router.Handle("/user", middleware.AuthMiddleware(http.HandlerFunc(s.UpdateUser))).Methods("PUT")                   // Update user profile

	logger.InfoLogger.Println("Registered routes:")
	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			logger.InfoLogger.Println("Route:", path)
		}
		return nil
	})
}
