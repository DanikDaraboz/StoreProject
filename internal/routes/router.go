package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// TODO middleware

	RegisterProductRoutes(router)
	RegisterOrderRoutes(router)
	RegisterHealthRoutes(router)

	return router
}
