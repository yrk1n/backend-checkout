package services

import (
	"github.com/yrk1n/backend-checkout/models"
	"github.com/yrk1n/backend-checkout/repositories"
)

type PromotionService struct {
	promotionRepo repositories.PromotionRepository
}

func NewPromotionService(promotionRepo repositories.PromotionRepository) *PromotionService {
	return &PromotionService{
		promotionRepo: promotionRepo,
	}
}

func (s *PromotionService) applySameSellerPromotion(cart *models.Cart, discount float64) {
	items := cart.Items
	if len(items) == 0 {
		return
	}

	firstSellerID := items[0].SellerId
	for _, item := range items {
		if item.SellerId != firstSellerID {
			return
		}
	}

	for _, item := range items {
		item.DiscountedPrice = item.Price * (1 - discount)
	}
}

func (s *PromotionService) applyCategoryPromotion(cart *models.Cart, discount float64) {
	for _, item := range cart.Items {
		if item.CategoryId == 3003 { // Assuming 3003 is the desired category ID for promotion
			item.DiscountedPrice = item.Price * (1 - discount)
		}
	}
}

func (s *PromotionService) GetApplicablePromotions(cart *models.Cart) ([]*models.Promotion, error) {
	return s.promotionRepo.GetApplicablePromotions(cart)
}

func (s *PromotionService) applyTotalPricePromotion(cart *models.Cart, discount float64) {
	// Logic to apply TotalPricePromotion discount to the cart
	// You can implement this based on your specific requirements
}

func (s *PromotionService) ApplyPromotions(cart *models.Cart) error {
	// Get applicable promotions for the cart
	applicablePromotions, err := s.promotionRepo.GetApplicablePromotions(cart)
	if err != nil {
		return err
	}

	// Apply each applicable promotion to the cart
	for _, promotion := range applicablePromotions {
		switch promotion.PromotionID {
		case 9909: // SameSellerPromotion
			s.applySameSellerPromotion(cart, promotion.DiscountValue)
		case 5676: // CategoryPromotion
			s.applyCategoryPromotion(cart, promotion.DiscountValue)
		case 1232: // TotalPricePromotion
			s.applyTotalPricePromotion(cart, promotion.DiscountValue)
			// Add more cases for other promotions if needed
		}
	}

	return nil
}
