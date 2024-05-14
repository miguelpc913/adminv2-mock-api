package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedAffiliateAgreementAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var affiliateAgreements []models.AffiliateAgreement
	var products []models.Product
	var buyerTypes []models.BuyerType
	db.Find(&buyerTypes)
	db.Find(&products)
	db.Find(&affiliateAgreements)
	for _, product := range products {
		for _, affiliateAgreement := range affiliateAgreements {
			db.Model(&affiliateAgreement).Association("ProductSet").Append(&product)
		}
	}

	for _, buyerType := range buyerTypes {
		for _, affiliateAgreement := range affiliateAgreements {
			db.Model(&affiliateAgreement).Association("BuyerTypeSet").Append(&buyerType)
		}
	}
}
