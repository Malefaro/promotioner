package enums

import "database/sql/driver"

type PromotionTypeEnum string

const (
	Discount                PromotionTypeEnum = "discount"
	Gift                    PromotionTypeEnum = "gift"
	CumulativeBonus         PromotionTypeEnum = "cumulative_bonus"
	SeveralProductsBuyBonus PromotionTypeEnum = "several_products_buy_bonus"
)

func (e *PromotionTypeEnum) Scan(value interface{}) error {
	*e = PromotionTypeEnum(value.([]byte))
	return nil
}

func (e PromotionTypeEnum) Value() (driver.Value, error) {
	return string(e), nil
}
