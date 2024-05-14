package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedPromotionPrices(db *gorm.DB) {
	promotionPrices := []models.PromotionPrice{
		{PromotionId: 1, BuyerTypeId: 1, SalesGroupId: 1, Percentage: 0.1},
		{PromotionId: 2, BuyerTypeId: 2, SalesGroupId: 2, Amount: 20},
		{PromotionId: 3, BuyerTypeId: 3, SalesGroupId: 3, Percentage: 0.1},
		{PromotionId: 4, BuyerTypeId: 4, SalesGroupId: 4, Amount: 40},
		{PromotionId: 5, BuyerTypeId: 5, SalesGroupId: 5, Percentage: 0.1},
	}

	result := db.Create(promotionPrices)
	if result.Error != nil {
		fmt.Println("Error seeding promotion_prices:", result.Error)
	}
}
