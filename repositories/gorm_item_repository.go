package repositories

import (
	"github.com/yrk1n/backend-checkout/database"
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
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *GORMItemRepository) Save(item *models.Item) error {
	return database.DB.Db.Save(item).Error
}

func (r *GORMItemRepository) Delete(id int) error {
	item := models.Item{ItemId: id}
	return database.DB.Db.Delete(&item).Error
}
func (r *GORMItemRepository) GetAll() ([]*models.Item, error) {
	var items []*models.Item
	if err := database.DB.Db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
