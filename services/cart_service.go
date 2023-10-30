package services

import (
	"errors"
	"fmt"

	"github.com/yrk1n/backend-checkout/database"
	"github.com/yrk1n/backend-checkout/models"
	"github.com/yrk1n/backend-checkout/repositories"
	"gorm.io/gorm"
)

type CartService struct {
	db          *gorm.DB
	cartRepo    repositories.CartRepository
	itemRepo    repositories.ItemRepository
	vasItemRepo repositories.VasItemRepository
}

func NewCartService(
	db *gorm.DB,
	cartRepo repositories.CartRepository,
	itemRepo repositories.ItemRepository,
	vasItemRepo repositories.VasItemRepository) *CartService {
	return &CartService{
		db:          db,
		cartRepo:    cartRepo,
		itemRepo:    itemRepo,
		vasItemRepo: vasItemRepo,
	}
}

func (s *CartService) InitializeCart() (*models.Cart, error) {
	var cart models.Cart
	result := s.db.First(&cart, 1)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		cart = models.Cart{CartID: 1}
		if result := s.db.Create(&cart); result.Error != nil {
			return nil, result.Error
		}
	}
	return &cart, nil
}

func (s *CartService) AddItemToCart(cartId int, item *models.Item) error {
	cart, err := s.cartRepo.GetByID(cartId)
	if err != nil {
		return err
	}
	cart.Items = append(cart.Items, item)
	return s.cartRepo.Save(cart)
}
func (s *CartService) GetItemByID(id int) (*models.Item, error) {
	var item models.Item
	err := database.DB.Db.Debug().Preload("VasItems").First(&item, id).Error
	return &item, err
}
func (s *CartService) GetVasItemRepo() repositories.VasItemRepository {
	return s.vasItemRepo
}

// func (s *CartService) AddVasItemToItem(cartId int, vasItem *models.VasItem) error {
// 	cart, err := s.cartRepo.GetByID(cartId)
// 	if err != nil {
// 		return err
// 	}

//		for i, cItem := range cart.Items {
//			if cItem.ItemId == vasItem.ParentItemId {
//				cart.Items[i].VasItems = append(cart.Items[i].VasItems, *vasItem)
//			}
//		}
//		return s.cartRepo.Save(cart)
//	}
func (s *CartService) AddVasItemToItem(vasItem *models.VasItem) error {
	item, err := s.itemRepo.GetByID(vasItem.ParentItemId)
	if err != nil {
		return err
	}

	for _, existingVasItem := range item.VasItems {
		if existingVasItem.VasItemId == vasItem.VasItemId {
			return fmt.Errorf("VasItem with ID %d already exists for the item", vasItem.VasItemId)
		}
	}

	err = s.itemRepo.AddVasItemToItem(vasItem, item)
	if err != nil {
		return err
	}

	return s.itemRepo.Save(item)
}

func (s *CartService) ResetCart(cartId int) error {
	cart, err := s.cartRepo.GetByID(cartId)
	if err != nil {
		return err
	}

	for _, item := range cart.Items {
		err := s.itemRepo.Delete(item.ItemId)
		if err != nil {
			return err
		}
	}

	// TODO: ADD OPTION TO DELETE VASITEMS
	return s.cartRepo.Save(cart)
}

func (s *CartService) GetCartByID(id int) (*models.Cart, error) {
	return s.cartRepo.GetByID(id)
}

// adding preloading to this fixed the issue took me a while to figure out
func (s *CartService) GetItems() ([]*models.Item, error) {
	var items []*models.Item
	err := s.db.Preload("VasItems").Find(&items).Error
	return items, err
}
