package repositories

import (
	"github.com/yrk1n/backend-checkout/models"
	"gorm.io/gorm"
)

type GORMCartRepository struct {
	db *gorm.DB
}

func NewGormCartRepository(db *gorm.DB) *GORMCartRepository {
	return &GORMCartRepository{db: db}
}

func (r *GORMCartRepository) AddItem(item *models.Item) error {
	var cart models.Cart
	if err := r.db.First(&cart).Error; err != nil {
		return err
	}
	cart.Items = append(cart.Items, item)
	return r.db.Save(&cart).Error
}

func (r *GORMCartRepository) RemoveItem(itemID int) error {
	var cart models.Cart
	if err := r.db.First(&cart).Error; err != nil {
		return err
	}
	for i, item := range cart.Items {
		if item.ItemId == itemID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}
	return r.db.Save(&cart).Error
}

func (r *GORMCartRepository) ResetCart() error {
	var cart models.Cart
	if err := r.db.First(&cart).Error; err != nil {
		return err
	}
	cart.Items = []*models.Item{}
	return r.db.Save(&cart).Error
}

func (r *GORMCartRepository) GetItems() ([]*models.Item, error) {
	var cart models.Cart
	if err := r.db.Preload("Items").First(&cart).Error; err != nil {
		return nil, err
	}
	return cart.Items, nil
}
func (r *GORMCartRepository) Save(cart *models.Cart) error {
	return r.db.Save(cart).Error
}

func (r *GORMCartRepository) GetByID(id int) (*models.Cart, error) {
	var cart models.Cart
	if err := r.db.Preload("Items").First(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}
