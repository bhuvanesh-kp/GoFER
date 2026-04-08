package main

import (
	internalHTTP "gofer/services/trip-service/internal/infrastructure/http"
	"gofer/services/trip-service/internal/infrastructure/repository"
	"gofer/services/trip-service/internal/service"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Respository
	inmemoryRepository := repository.NewInmemeoryRespository()

	// Service
	svc := service.NewService(inmemoryRepository)

	httphandler := internalHTTP.HttpHandler{
		Service: svc,
	}

	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
