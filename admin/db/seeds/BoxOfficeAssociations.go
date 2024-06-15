package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedBoxOfficeAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var boxOffices []models.BoxOffice
	var salesGroups []models.SalesGroup
	var products []models.Product
	db.Find(&boxOffices)
	db.Find(&products)
	db.Find(&salesGroups)
	for _, salesGroup := range salesGroups {
		for _, boxOffice := range boxOffices {
			db.Model(&boxOffice).Association("SalesGroupSet").Append(&salesGroup)
		}
	}
	for _, product := range products {
		for _, boxOffice := range boxOffices {
			db.Model(&boxOffice).Association("ProductSet").Append(&product)
		}
	}
}
