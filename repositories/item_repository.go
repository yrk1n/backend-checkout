package repositories

import "github.com/yrk1n/backend-checkout/models"

type ItemRepository interface {
	GetByID(id int) (*models.Item, error)
	Save(item *models.Item) error
	Delete(id int) error
	GetAll() ([]*models.Item, error)
	AddVasItemToItem(vasItem *models.VasItem, item *models.Item) error
}
