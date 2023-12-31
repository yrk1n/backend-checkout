package repositories

import (
	"fmt"

	"github.com/yrk1n/backend-checkout/models"
	"gorm.io/gorm"
)

type GORMVasItemRepository struct {
	db *gorm.DB
}

func NewGORMVasItemRepository(db *gorm.DB) *GORMVasItemRepository {
	return &GORMVasItemRepository{db: db}
}

func (r *GORMVasItemRepository) Create(vasItem *models.VasItem) error {
	return r.db.Create(vasItem).Error
}
func (r *GORMVasItemRepository) GetByID(id int) (*models.VasItem, error) {
	var vasItem models.VasItem
	if err := r.db.First(&vasItem, id).Error; err != nil {
		return nil, err
	}
	return &vasItem, nil
}

func (r *GORMVasItemRepository) Save(vasItem *models.VasItem) error {
	return r.db.Save(vasItem).Error
}

func (r *GORMVasItemRepository) Delete(id int) error {
	vasItem := models.VasItem{VasItemId: id}
	return r.db.Delete(&vasItem).Error
}

func (r *GORMVasItemRepository) GetAll() ([]*models.VasItem, error) {
	var vasItems []*models.VasItem
	if err := r.db.Find(&vasItems).Error; err != nil {
		return nil, err
	}
	return vasItems, nil
}
func (r *GORMVasItemRepository) AddVasItemToItem(vasItem *models.VasItem, item *models.Item) error {
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

func (r *GORMVasItemRepository) CreateVasItemForItem(itemID int, vasItem *models.VasItem) error {
	var item models.Item
	if err := r.db.First(&item, itemID).Error; err != nil {
		return err
	}

	vasItem.ParentItemId = item.ItemId

	if err := r.db.Create(vasItem).Error; err != nil {
		return err
	}

	return nil
}
