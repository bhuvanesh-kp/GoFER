package main

import (
	"context"
	internalHTTP "gofer/services/trip-service/internal/infrastructure/http"
	"gofer/services/trip-service/internal/infrastructure/repository"
	"gofer/services/trip-service/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	serverError := make(chan error, 1)

	go func() {
		log.Printf("Server listening on %s", server.Addr)
		serverError <- server.ListenAndServe()
	}()

	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Printf("Error starting the server: %v", err)
	case sig := <-shutDown:
		log.Printf("Server is shutting down due to %v signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop the server gracefully: %v", err)
			server.Close()
		}
	}

	// TODO: after finishing project added request based graceful shutdown rather than fixed value shutdown
}
