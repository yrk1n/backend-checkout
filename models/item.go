package models

type ItemType string

const (
	Digital ItemType = "Digital"
	Default ItemType = "Default"
)

type Item struct {
	ItemId int `json:"item_id" gorm:"primaryKey"`

	CartID     int       `gorm:"column:cart_id"`
	CategoryId int       `json:"category"`
	SellerId   int       `json:"seller_id"`
	Price      float64   `json:"price"`
	Quantity   int       `json:"quantity"`
	Type       ItemType  `json:"type"`
	VasItems   []VasItem `json:"vas_items" gorm:"foreignKey:ParentItemId;onDelete:CASCADE"`
}
