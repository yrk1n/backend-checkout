package models

type VasItem struct {
	Item
	VasItemId    int `json:"vas_item_id" gorm:"primaryKey"`
	ParentItemId int `json:"parent_item_id" gorm:"foreignKey:ParentItemId;references:ItemId"`
}
