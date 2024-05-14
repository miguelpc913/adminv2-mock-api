package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedPromotionAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var promotions []models.Promotion
	var products []models.Product
	var salesGroups []models.SalesGroup
	var buyerTypes []models.BuyerType
	db.Find(&promotions)
	db.Find(&products)
	db.Find(&salesGroups)
	db.Find(&buyerTypes)
	// Seed associations
	for _, product := range products {
		for _, promotion := range promotions {
			db.Model(&promotion).Association("ProductSet").Append(&product)
		}
	}
	for _, salesGroup := range salesGroups {
		for _, promotion := range promotions {
			db.Model(&promotion).Association("SalesGroupSet").Append(&salesGroup)
		}
	}
	for _, buyerType := range buyerTypes {
		for _, promotion := range promotions {
			db.Model(&promotion).Association("BuyerTypeSet").Append(&buyerType)
		}
	}
}
