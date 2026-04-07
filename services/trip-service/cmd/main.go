package main

import (
	"context"
	"gofer/services/trip-service/internal/domain"
	"gofer/services/trip-service/internal/infrastructure/repository"
	"gofer/services/trip-service/internal/service"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	sampel_fare := domain.RideFareModel{
		UserID: "17",
	}

	inmemory := repository.NewInmemeoryRespository()

	svc := service.NewService(inmemory)
	res, err := svc.CreateTrip(ctx, &sampel_fare)

	if err != nil{
		log.Fatal(err)
	}

	log.Println(res)

	// temporarily keeping server up (never do this in real world application - only for quick test)
	for {
		time.Sleep(time.Second * 1)
	}
}
