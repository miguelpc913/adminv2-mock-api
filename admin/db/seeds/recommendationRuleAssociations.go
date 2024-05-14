package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedRecommendationRulesAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var recommendationRules []models.RecommendationRule
	var salesGroups []models.SalesGroup
	var buyerTypes []models.BuyerType
	db.Find(&salesGroups)
	db.Find(&buyerTypes)
	db.Find(&recommendationRules)
	for _, salesGroup := range salesGroups {
		for _, recommendationRule := range recommendationRules {
			db.Model(&recommendationRule).Association("SalesGroupSet").Append(&salesGroup)
		}
	}
	for _, buyerType := range buyerTypes {
		for _, recommendationRule := range recommendationRules {
			db.Model(&recommendationRule).Association("BuyerTypeSet").Append(&buyerType)
		}
	}
}
