package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedProductInfoAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var productsInfo []models.ProductInfo
	var salesGroups []models.SalesGroup
	var products []models.Product
	var venues []models.Venue
	db.Find(&products)
	db.Find(&salesGroups)
	db.Find(&venues)
	db.Find(&productsInfo)
	for _, salesGroup := range salesGroups {
		for _, productInfo := range productsInfo {
			db.Model(&productInfo).Association("SalesGroupSet").Append(&salesGroup)
		}
	}
	for _, product := range products {
		for _, productInfo := range productsInfo {
			db.Model(&productInfo).Association("ProductSet").Append(&product)
		}
	}
	for _, venue := range venues {
		for _, productInfo := range productsInfo {
			db.Model(&productInfo).Association("VenueSet").Append(&venue)
		}
	}
}
