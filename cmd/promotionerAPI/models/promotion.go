package models

import "time"

type PromotionType struct {
	PromotionTypeID int64
	Name            string
}

type Promotion struct {
	PromotionID   int64
	StartDate     *time.Time
	EndDate       *time.Time
	Text          string
	Banner        string
	Rating        float64 // 0-10 stars
	Addresses     []int64
	IsUserCreated bool
	// foreign keys
	PromotionTypeID int64
	FoodPlaceID     int64
}
