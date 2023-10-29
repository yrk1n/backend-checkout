package repositories

import (
	"github.com/yrk1n/backend-checkout/models"

	"gorm.io/gorm"
)

type GORMVasItemRepository struct {
	db *gorm.DB
}

func NewGORMVasItemRepository(db *gorm.DB) *GORMVasItemRepository {
	return &GORMVasItemRepository{db: db}
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
