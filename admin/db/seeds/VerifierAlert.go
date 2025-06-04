package seeds

import (
	"fmt"
	"time"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedVerifierAlerts(db *gorm.DB) {
	// VerifierAlertPromotion entries with valid sounds and hex alert colors
	promotionAlerts := []models.VerifierAlertPromotion{
		{
			PromotionID: 1,

			AlertColor: "#FF0000", // red
			AlertSound: "train_horn",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			PromotionID: 2,

			AlertColor: "#00FF00", // green
			AlertSound: "bell",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	if err := db.Create(&promotionAlerts).Error; err != nil {
		fmt.Println("❌ Error seeding VerifierAlertPromotion:", err)
	}

	// VerifierAlertBuyerType entries with valid sounds and hex colors
	buyerTypeAlerts := []models.VerifierAlertBuyerType{
		{
			BuyerTypeID: 1,

			AlertColor: "#FFD700", // gold
			AlertSound: "aplause",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			BuyerTypeID: 2,

			AlertColor: "#1E90FF", // blue
			AlertSound: "standard_sound_1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	if err := db.Create(&buyerTypeAlerts).Error; err != nil {
		fmt.Println("❌ Error seeding VerifierAlertBuyerType:", err)
	}
}
