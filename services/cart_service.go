package services

import (
	"errors"
	"fmt"

	"github.com/yrk1n/backend-checkout/database"
	"github.com/yrk1n/backend-checkout/models"
	"github.com/yrk1n/backend-checkout/repositories"
	"gorm.io/gorm"
)

const (
	maxUniqueItems        = 10
	maxTotalItems         = 30
	maxCartAmount         = 500000
	digitalCategoryID     = 7889
	maxDigitalItemQty     = 5
	furnitureCategoryID   = 1001
	electronicsCategoryID = 3004
	vasCategoryID         = 3242
	vasSellerID           = 5003
	maxVASForItem         = 3
)

type CartService struct {
	db          *gorm.DB
	cartRepo    repositories.CartRepository
	itemRepo    repositories.ItemRepository
	vasItemRepo repositories.VasItemRepository
	promoSvc    *PromotionService
}

func NewCartService(
	db *gorm.DB,
	cartRepo repositories.CartRepository,
	itemRepo repositories.ItemRepository,
	promoSvc *PromotionService,
	vasItemRepo repositories.VasItemRepository) *CartService {
	return &CartService{
		db:          db,
		cartRepo:    cartRepo,
		itemRepo:    itemRepo,
		vasItemRepo: vasItemRepo,
		promoSvc:    promoSvc,
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

	if len(cart.Items) >= maxUniqueItems {
		return fmt.Errorf("cart already contains %d unique items", maxUniqueItems)
	}

	totalItems := 0
	totalPrice := 0.0
	hasDigital := false

	for _, cItem := range cart.Items {
		totalItems += cItem.Quantity
		totalPrice += cItem.Price * float64(cItem.Quantity)

		if cItem.CategoryId == digitalCategoryID {
			hasDigital = true
		}
	}

	if hasDigital && item.CategoryId != digitalCategoryID {
		return errors.New("mixed item types in cart")
	}

	if totalItems+item.Quantity > maxTotalItems {
		return errors.New("maximum total items in cart exceeded")
	}

	if totalPrice+item.Price*float64(item.Quantity) > maxCartAmount {
		return errors.New("maximum cart amount exceeded")
	}

	cart.Items = append(cart.Items, item)

	// Apply promotions if available
	if s.promoSvc != nil {
		if err := s.promoSvc.ApplyPromotions(cart); err != nil {
			return err
		}
	}

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

func (s *CartService) ResetCart() error {
	hardCodedCartID := 1
	return s.db.Where("cart_id = ?", hardCodedCartID).Delete(&models.Item{}).Error
}

func (s *CartService) GetCartByID(id int) (*models.Cart, error) {
	return s.cartRepo.GetByID(id)
}

// adding preloading to this according to fixed the issue took me a while to figure out
func (s *CartService) GetItems() ([]*models.Item, error) {
	var items []*models.Item
	err := s.db.Preload("VasItems").Find(&items).Error
	return items, err
}

func (s *CartService) ValidateCart(cart *models.Cart) error {
	if len(cart.Items) > 10 {
		return errors.New("cart can contain a maximum of 10 unique items")
	}

	totalProducts := 0
	for _, item := range cart.Items {
		totalProducts += item.Quantity
	}
	if totalProducts > 30 {
		return errors.New("the total number of products cannot exceed 30")
	}

	if cart.TotalPrice > 500000 {
		return errors.New("the total amount of the cart cannot exceed 500,000 TL")
	}

	return nil
}

func (s *CartService) ValidateItem(item *models.Item) error {
	if item.Quantity > 10 {
		return errors.New("maximum quantity for an item is 10")
	}

	if item.CategoryId == 7889 && item.Quantity > 5 {
		return errors.New("maximum quantity for a digital item is 5")
	}

	return nil
}
