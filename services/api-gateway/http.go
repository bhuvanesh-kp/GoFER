package main

import (
	"encoding/json"
	"gofer/services/api-gateway/grpc_clients"
	"gofer/shared/contracts"
	"log"
	"net/http"
)

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
		log.Fatal(err)
	}

	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.toProto())
	if err != nil {
		log.Fatal(err)
		return
	}

	reqRes := contracts.APIResponse{
		Data: tripPreview,
	}

	writeJSON(w, http.StatusAccepted, reqRes)
}
