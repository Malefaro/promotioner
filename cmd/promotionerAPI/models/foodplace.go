package models

import "promotioner/cmd/promotionerAPI/models/enums"

type FoodPlace struct {
	FoodPlaceID   int64  `json:"food_place_id"`
	Name          string `json:"name"`
	PriceType     int16  `json:"price_type"`
	Description   string `json:"description"`
	FoodPlaceType enums.FoodPlaceType
	// TODO: other fields
}

type FoodPlaceAddress struct {
	FoodPlaceAddressID int64   `json:"food_place_address_id"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	// foreign keys
	FoodPlaceID int64 `json:"food_place_id"`
}
