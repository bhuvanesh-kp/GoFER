package main

import (
	"encoding/json"
	"gofer/shared/contracts"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Invalid Request", http.StatusBadGateway)
		return
	}

	var reqBody previewTripRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "failed to parse json data", http.StatusBadRequest)
		return
	}

	if reqBody.UserId == ""{
		http.Error(w, "UserId is required", http.StatusBadRequest)
		return
	}

	reqRes := contracts.APIResponse{
		Data: "ok",
	}

	writeJSON(w, http.StatusAccepted, reqRes)
}