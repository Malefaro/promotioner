package models

type FoodPlace struct {
	FoodPlaceID int64
	Name        string
	PriceType   int32
	// TODO: other fields
}

type FoodPlaceAddress struct {
	FoodPlaceAddressID int64
	Latitude           float64
	Longitude          float64

	FoodPlaceID int64
}
