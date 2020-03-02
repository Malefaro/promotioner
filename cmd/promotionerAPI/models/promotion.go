package models

import (
	"promotioner/cmd/promotionerAPI/models/enums"
	"time"
)

type PromotionCreator struct {
	PromotionCreatorID int64 `json:"promotion_creator_id"`
	IsCustom           bool  `json:"is_custom"`
	// foreign keys
	UserID      int64 `json:"user_id"`
	PromotionID int64 `json:"promotion_id"`
}

type Promotion struct {
	PromotionID   int64                   `json:"promotion_id"`
	StartDate     *time.Time              `json:"start_date"`
	EndDate       *time.Time              `json:"end_date"`
	LabelText     string                  `json:"label_text"`
	Description   string                  `json:"description"`
	Banner        string                  `json:"banner"`
	Rating        float32                 `json:"rating"` // 0-10 calculated rating
	PromotionType enums.PromotionTypeEnum `json:"promotion_type"`
	// foreign keys
	CreatorID   int64 `json:"creator_id"`
	FoodPlaceID int64 `json:"food_place_id"`
}

// ManyToMany
type PromotionAddress struct {
	PromotionAddressID int64 `json:"promotion_address_id"`
	// foreign keys
	PromotionID        int64 `json:"promotion_id"`
	FoodPlaceAddressID int64 `json:"food_place_address_id"`
}

type PromotionLike struct {
	PromotionLikeID int64 `json:"promotion_like_id"`
	Rating          int16 `json:"rating"` // 0-10 stars
	// foreign keys
	UserID      int64 `json:"user_id"`
	PromotionID int64 `json:"promotion_id"`
}
