package enums

import "database/sql/driver"

type FoodPlaceType string

const (
	Cafe FoodPlaceType = "cafe"
	Bar FoodPlaceType = "bar"
	Restaurant FoodPlaceType = "restaurant"
	Club FoodPlaceType = "club"
)

func (e *FoodPlaceType) Scan(value interface{}) error {
	*e = FoodPlaceType(value.([]byte))
	return nil
}

func (e FoodPlaceType) Value() (driver.Value, error) {
	return string(e), nil
}
