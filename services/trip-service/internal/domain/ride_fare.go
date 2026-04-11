package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // suv, sedan, van, luxury type ride
	TotalPriceInCents float64
}
