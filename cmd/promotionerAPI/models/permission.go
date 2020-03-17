package models

type Permission struct {
	PermissionID int64  `json:"permission_id"`
	Name         string `json:"name"`
	Priority     int16  `json:"priority"`
}

type PermissionGroup struct {
	PermissionGroupID int64  `json:"permission_group_id"`
	Name              string `json:"name"`
}

// ManyToMany
type PermissionGroupPermission struct {
	PermissionGroupPermissionID int64 `json:"permission_group_permission_id"`

	PermissionID      int64 `json:"permission_id"`
	PermissionGroupID int64 `json:"permission_group_id"`
}

type FoodPlacePermissionGroup struct {
	FoodPlacePermissionGroupID int64 `json:"food_place_permission_group_id"`

	FoodPlaceID       int64 `json:"food_place_id"`
	PermissionGroupID int64 `json:"permission_group_id"`
}

type UserFoodPlacePermissionGroup struct {
	UserFoodPlacePermissionGroupID int64 `json:"user_food_place_permission_group_id"`

	FoodPlacePermissionGroupID int64 `json:"food_place_permission_group_id"`
	UserID                     int64 `json:"user_id"`
}
