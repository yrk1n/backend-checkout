package repositories

import (
	"github.com/yrk1n/backend-checkout/models"
)

type CartRepository interface {
	AddItem(item *models.Item) error
	RemoveItem(itemID int) error
	ResetCart() error
	GetItems() ([]*models.Item, error)
}
