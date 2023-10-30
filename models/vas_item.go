package models

type VasItem struct {
	VasItemId      int     `json:"vas_item_id" gorm:"primaryKey"`
	ParentItemId   int     `json:"parent_item_id" gorm:"foreignKey:ParentItemId;references:ItemId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	ActivationCode string  `json:"activation_code"`
}
