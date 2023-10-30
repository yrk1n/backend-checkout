package repositories

import (
	"fmt"

	"github.com/yrk1n/backend-checkout/models"
	"gorm.io/gorm"
)

type GORMItemRepository struct {
	db *gorm.DB
}

func NewGORMItemRepository(db *gorm.DB) *GORMItemRepository {
	return &GORMItemRepository{db: db}
}

func (r *GORMItemRepository) GetByID(id int) (*models.Item, error) {
	var item models.Item
	if err := r.db.Preload("VasItems").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *GORMItemRepository) Save(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *GORMItemRepository) Delete(id int) error {
	item := models.Item{ItemId: id}
	return r.db.Delete(&item).Error
}
func (r *GORMItemRepository) GetAll() ([]*models.Item, error) {
	var items []*models.Item
	if err := r.db.Preload("VasItems").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *GORMItemRepository) AddVasItemToItem(vasItem *models.VasItem, item *models.Item) error {
	var existingItem models.Item
	if err := r.db.Preload("VasItems").First(&existingItem, item.ItemId).Error; err != nil {
		return err
	}

	association := r.db.Model(&existingItem).Association("VasItems")
	if err := association.Append(vasItem); err != nil {
		return fmt.Errorf("failed to associate vas item: %v", err)
	}
	return nil
}
