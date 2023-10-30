package models

type PromotionType string

const (
	SameSeller PromotionType = "SameSeller"
	Category   PromotionType = "Category"
	TotalPrice PromotionType = "TotalPrice"
)

type Promotion struct {
	PromotionID   int           `json:"promotion_id" gorm:"primaryKey"`
	Type          PromotionType `json:"type"`
	CategoryID    *int          `json:"category_id"` // TODO: Clarify this since Category model is removed
	DiscountValue float64       `json:"discount_value"`
	MinTotalPrice *float64      `json:"min_total_price"`
	MaxTotalPrice *float64      `json:"max_total_price"`
}
