package repositories

import (
	"github.com/yrk1n/backend-checkout/models"
	"gorm.io/gorm"
)

type GORMPromotionRepository struct {
	db *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &GORMPromotionRepository{db: db}
}

func (r *GORMPromotionRepository) GetApplicablePromotions(cart *models.Cart) ([]*models.Promotion, error) {
	var applicablePromotions []*models.Promotion

	// Check for SameSellerPromotion
	if r.isSameSeller(cart) {
		applicablePromotions = append(applicablePromotions, &models.Promotion{PromotionID: 9909, DiscountValue: 0.10}) // 10% discount
	}

	// Check for CategoryPromotion
	if r.hasCategoryID(cart, 3003) {
		applicablePromotions = append(applicablePromotions, &models.Promotion{PromotionID: 5676, DiscountValue: 0.05}) // 5% discount
	}

	// Check for TotalPricePromotion
	totalPriceDiscount := r.getTotalPriceDiscount(cart.TotalPrice)
	if totalPriceDiscount > 0 {
		applicablePromotions = append(applicablePromotions, &models.Promotion{PromotionID: 1232, DiscountValue: totalPriceDiscount})
	}

	return applicablePromotions, nil
}

func (r *GORMPromotionRepository) isSameSeller(cart *models.Cart) bool {
	if len(cart.Items) == 0 {
		return false
	}
	firstSellerID := cart.Items[0].SellerId
	for _, item := range cart.Items {
		if item.SellerId != firstSellerID || item.Type == "VasItem" {
			return false
		}
	}
	return true
}

func (r *GORMPromotionRepository) hasCategoryID(cart *models.Cart, categoryID int) bool {
	for _, item := range cart.Items {
		if item.CategoryId == categoryID {
			return true
		}
	}
	return false
}

func (r *GORMPromotionRepository) getTotalPriceDiscount(totalPrice float64) float64 {
	switch {
	case totalPrice < 5000:
		return 250
	case totalPrice >= 5000 && totalPrice < 10000:
		return 500
	case totalPrice >= 10000 && totalPrice < 50000:
		return 1000
	case totalPrice >= 50000:
		return 2000
	}
	return 0
}

func (r *GORMPromotionRepository) applySameSellerPromotion(cart *models.Cart, discount float64) {
	for _, item := range cart.Items {
		item.Price = item.Price * (1 - discount)
	}
}

func (r *GORMPromotionRepository) applyCategoryPromotion(cart *models.Cart, discount float64) {
	for _, item := range cart.Items {
		if item.CategoryId == 3003 {
			item.Price = item.Price * (1 - discount)
		}
	}
}

func (r *GORMPromotionRepository) applyTotalPricePromotion(cart *models.Cart, discount float64) {
	cart.TotalPrice = cart.TotalPrice * (1 - discount)
}
