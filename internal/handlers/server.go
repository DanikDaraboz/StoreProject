package handlers

import (
	"html/template"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/gorilla/mux"
)

type Server struct {
	Router         *mux.Router
	DB             *mongoDriver.Client
	Services       *services.Services
	TemplatesCache map[string]*template.Template
}

func NewServer(router *mux.Router, dbConn *mongoDriver.Client, services *services.Services, templates map[string]*template.Template) *Server {
	s := &Server{
		Router:         router,
		DB:             dbConn,
		Services:       services,
		TemplatesCache: templates,
	}

	return s
}
