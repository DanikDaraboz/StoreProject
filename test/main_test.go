package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DanikDaraboz/StoreProject/internal/handlers"
	"github.com/DanikDaraboz/StoreProject/internal/repository"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	"github.com/DanikDaraboz/StoreProject/internal/routes"
	"github.com/DanikDaraboz/StoreProject/internal/services"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
	"github.com/gorilla/mux"
)

var srv *handlers.Server 

func TestMain(m *testing.M) {
	router := mux.NewRouter()

	mongoURI := "mongodb://localhost:27017"
	mongoClient, err := mongo.Connect(mongoURI)
	if err != nil {
		logger.ErrorLogger.Println("Failed to connect to mongo:", err)
	}

	testDB := mongoClient.Database("ecommerce_test")
	testRepos := repository.NewRepositories(testDB)
	testServices := services.NewServices(testRepos)

	templateCache, err := handlers.NewTemplateCache()
	if err != nil {
		panic("Failed to load templates: " + err.Error())
	}

	srv = handlers.NewServer(router, mongoClient, testServices, templateCache)

	routes.RegisterRoutes(srv)

	exitCode := m.Run()

	mongoClient.Disconnect(context.Background()) 
	os.Exit(exitCode)
}

// Helper function to execute test requests
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	srv.Router.ServeHTTP(rr, req)
	return rr
}

// Helper to check response status
func checkStatus(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected status %d but got %d", expected, actual)
	}
}
