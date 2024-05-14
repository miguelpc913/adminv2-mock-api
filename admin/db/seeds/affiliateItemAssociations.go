package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedAffiliateItemsAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var affiliateItems []models.AffiliateItem
	var products []models.Product
	db.Find(&products)
	db.Find(&affiliateItems)
	for _, product := range products {
		for _, affiliateItem := range affiliateItems {
			db.Model(&affiliateItem).Association("ProductSet").Append(&product)
		}
	}
}
