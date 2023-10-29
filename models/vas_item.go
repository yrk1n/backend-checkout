package models

type VasItem struct {
	VasItemId    int     `json:"vas_item_id" gorm:"primaryKey"`
	ParentItemId int     `json:"parent_item_id" gorm:"foreignKey:ItemId"`
	CategoryId   int     `json:"category_id"`
	SellerId     int     `json:"seller_id"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	CartID       int     `json:"cart_id"`
}