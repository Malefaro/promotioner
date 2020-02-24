package models

import "time"

type PromotionType struct {
	PromotionTypeID int
	Name string
}

type Promotion struct {
	PromotionID int
	StartDate time.Time
	EndDate time.Time
	Text string
	Banner string
	// foreign keys
	PromotionTypeID int
	FoodPlaceID int
}
