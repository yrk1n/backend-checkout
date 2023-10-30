package repositories

import (
	"github.com/yrk1n/backend-checkout/models"
)

type PromotionRepository interface {
	GetApplicablePromotions(cart *models.Cart) ([]*models.Promotion, error)
	applySameSellerPromotion(cart *models.Cart, discount float64)
	applyCategoryPromotion(cart *models.Cart, discount float64)
	applyTotalPricePromotion(cart *models.Cart, discount float64)
}
