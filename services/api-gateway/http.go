package main

import (
	"encoding/json"
	"gofer/services/api-gateway/grpc_clients"
	"gofer/shared/contracts"
	"log"
	"net/http"
)

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	var reqBody startTripPreview
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	defer tripService.Close()

	trip, err := tripService.Client.CreateTrip(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("Failed to start a trip: %v", err)
		http.Error(w, "Failed to start trip", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: trip}

	writeJSON(w, http.StatusCreated, response)
}

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusBadGateway)
		return
	}

	var reqBody previewTripRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "failed to parse json data", http.StatusBadRequest)
		log.Println("Error @ api-gatway")
		return
	}

	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "UserId is required", http.StatusBadRequest)
		return
	}

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Printf("failed to create trip service client: %v", err)
		http.Error(w, "trip service unavailable", http.StatusServiceUnavailable)
		return
	}

	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.toProto())
	if err != nil {
		log.Printf("PreviewTrip failed: %v", err)
		http.Error(w, "failed to preview trip", http.StatusBadGateway)
		return
	}

	reqRes := contracts.APIResponse{
		Data: tripPreview,
	}

	writeJSON(w, http.StatusAccepted, reqRes)
}
