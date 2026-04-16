package main

import (
	"gofer/shared/types"
	pb "gofer/shared/proto/trip"
)

type previewTripRequest struct {
	UserId      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (p *previewTripRequest) toProto() *pb.PreviewTripRequest {
	return &pb.PreviewTripRequest{
		UserID: p.UserId,
		StartLocation: &pb.Coordinate{
			Latitude:  p.PickUp.Latitude,
			Longitude: p.PickUp.Longitude,
		},
		EndLocation: &pb.Coordinate{
			Latitude:  p.Destination.Latitude,
			Longitude: p.Destination.Longitude,
		},
	}
}
