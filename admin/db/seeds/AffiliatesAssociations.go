package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedAffiliateAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var affiliates []models.Affiliate
	var affiliateAgreements []models.AffiliateAgreement
	db.Find(&affiliateAgreements)
	db.Find(&affiliates)
	for _, affiliateAgreement := range affiliateAgreements {
		for _, affiliate := range affiliates {
			db.Model(&affiliate).Association("AffiliateAgreementSet").Append(&affiliateAgreement)
		}
	}
}
