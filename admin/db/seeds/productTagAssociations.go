package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedProductTagAssociation(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var products []models.Product
	var productTags []models.ProductTag
	var salesGroups []models.SalesGroup
	db.Find(&products)
	db.Find(&productTags)
	db.Find(&salesGroups)
	// Seed associations
	for _, product := range products {
		for _, productTag := range productTags {
			db.Model(&productTag).Association("ProductSet").Append(&product)
		}
	}
	for _, salesGroup := range salesGroups {
		for _, productTag := range productTags {
			db.Model(&productTag).Association("SalesGroupSet").Append(&salesGroup)
		}
	}
}
