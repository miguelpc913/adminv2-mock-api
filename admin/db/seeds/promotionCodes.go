package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedPromotionCodes(db *gorm.DB) {
	promotionPrices := []models.PromotionalCode{
		{PromotionId: 2, Code: "ABC123", Quantity: 100},
		{PromotionId: 2, Code: "DEF456", Quantity: 150},
		{PromotionId: 2, Code: "GHI789", Quantity: 200},
		{PromotionId: 2, Code: "JKL012", Quantity: 250},
		{PromotionId: 2, Code: "MNO345", Quantity: 300},
	}

	result := db.Create(promotionPrices)
	if result.Error != nil {
		fmt.Println("Error seeding promotion_prices:", result.Error)
	}
}
