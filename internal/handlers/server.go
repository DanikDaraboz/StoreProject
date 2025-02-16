package handlers

import (
	"html/template"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/DanikDaraboz/StoreProject/pkg/middleware"
	"github.com/gorilla/mux"
)

type Server struct {
	Router         *mux.Router
	DB             *mongoDriver.Client
	Services       *services.Services
	TemplatesCache map[string]*template.Template
	Middleware     middleware.MiddlewareInterface
}

func NewServer(router *mux.Router, dbConn *mongoDriver.Client, services *services.Services, templates map[string]*template.Template, middlewares middleware.MiddlewareInterface) *Server {
	return &Server{
		Router:         router,
		DB:             dbConn,
		Services:       services,
		TemplatesCache: templates,
		Middleware:     middlewares,
	}
}
