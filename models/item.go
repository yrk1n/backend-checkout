package models

type ItemType string

const (
	Digital ItemType = "Digital"
	Default ItemType = "Default"
)

type Item struct {
	ItemId      int        `json:"item_id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Quantity    int        `json:"quantity"`
	SellerId    int        `json:"seller_id"`
	CategoryId  int        `json:"category_id"`
	CartID      int        `json:"cart_id"`
	VasItems    []*VasItem `json:"vas_items" gorm:"foreignKey:ParentItemId"`
}
