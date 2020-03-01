package models

import (
	"promotioner/cmd/promotionerAPI/models/enums"
	"time"
)

type PromotionCreator struct {
	PromotionCreatorID int64
	IsCustom           bool
	// foreign keys
	UserID      int64
	PromotionID int64
}

type Promotion struct {
	PromotionID   int64
	StartDate     *time.Time
	EndDate       *time.Time
	LabelText     string
	Description   string
	Banner        string
	Rating        float32 // 0-10 calculated rating
	PromotionType enums.PromotionTypeEnum
	// foreign keys
	CreatorID   int64
	FoodPlaceID int64
}

// ManyToMany
type PromotionAddressRelation struct {
	PromotionAddressRelationID int64
	// foreign keys
	PromotionID        int64
	FoodPlaceAddressID int64
}

type PromotionLike struct {
	PromotionLikeID int64
	Rating          int16 // 0-10 stars
	// foreign keys
	UserID      int64
	PromotionID int64
}
