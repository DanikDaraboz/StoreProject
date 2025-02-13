package handlers

import (
	"context"
	"net/http"
	"time"
)

func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := s.DB.Ping(ctx, nil) 
	if err != nil {
		http.Error(w, "MongoDB not connected", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("MongoDB is connected!"))
}
