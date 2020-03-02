package models

import "promotioner/cmd/promotionerAPI/models/enums"

type FoodPlace struct {
	FoodPlaceID   int64               `json:"food_place_id"`
	Name          string              `json:"name"`
	AveragePrice  int64               `json:"price_type"`
	Description   string              `json:"description"`
	FoodPlaceType enums.FoodPlaceType `json:"food_place_type"`
	FoodType      enums.FoodType      `json:"food_type"`
	Rating        float32             `json:"rating"`
}

type FoodPlaceAddress struct {
	FoodPlaceAddressID int64   `json:"food_place_address_id"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Address            string  `json:"address"`
	// foreign keys
	FoodPlaceID int64 `json:"food_place_id"`
}

// ManyToMany
type FoodPlaceLike struct {
	FoodPlaceLikeID int64 `json:"food_place_like_id"`
	Rating          int16 `json:"rating"`
	// foreign keys
	FoodPlaceID int64 `json:"food_place_id"`
	UserID      int64 `json:"user_id"`
}
