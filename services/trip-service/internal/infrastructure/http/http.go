package http

import (
	"encoding/json"
	"gofer/services/trip-service/internal/domain"
	"gofer/shared/types"
	"log"
	"net/http"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserId      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		log.Println("Error from trip-service/HandleTripPreview")
		return
	}

	ctx := r.Context()

	trip, err := s.Service.GetRoute(ctx, &reqBody.PickUp, &reqBody.Destination)
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, http.StatusOK, trip)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
