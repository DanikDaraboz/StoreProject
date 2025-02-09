package handlers

import (
	"net/http"

	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	err := mongo.PingMongoDB()
	if err != nil {
		http.Error(w, "MongoDB not connected", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MongoDB is connected!"))
}
