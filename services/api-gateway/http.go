package main

import (
	"bytes"
	"encoding/json"
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

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed to parse JSON data from trip service", http.StatusBadRequest)
		return
	}

	reqRes := contracts.APIResponse{
		Data: respBody,
	}

	writeJSON(w, http.StatusAccepted, reqRes)
}
