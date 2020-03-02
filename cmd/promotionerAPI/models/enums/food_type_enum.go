package enums

import "database/sql/driver"

type FoodType string

const (
	Sushi   FoodType = "sushi"
	Pizza   FoodType = "pizza"
	Burgers FoodType = "burgers"
	Italian FoodType = "italian"
	Asian   FoodType = "asian"
	Eastern FoodType = "eastern"
	Russian FoodType = "russian"
)

func (e *FoodType) Scan(value interface{}) error {
	*e = FoodType(value.([]byte))
	return nil
}

func (e FoodType) Value() (driver.Value, error) {
	return string(e), nil
}
