package models

type Cart struct {
	CartID         int     `json:"cart_id" gorm:"primaryKey"`
	MaxUniqueItems int     `json:"max_unique_items"`
	MaxTotalItems  int     `json:"max_total_items"`
	MaxTotalPrice  int     `json:"max_total_price"`
	TotalPrice     float64 `json:"total_price"`
	Items          []*Item `gorm:"foreignKey:CartID;references:CartID"`
}
