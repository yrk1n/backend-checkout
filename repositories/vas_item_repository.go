package repositories

import (
	"github.com/yrk1n/backend-checkout/models"
)

type VasItemRepository interface {
	GetByID(id int) (*models.VasItem, error)
	Save(item *models.VasItem) error
	Delete(id int) error
	GetAll() ([]*models.VasItem, error)
}
