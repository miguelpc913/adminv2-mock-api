package seeds

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedSalesGroupsAssociations(db *gorm.DB) {
	// Retrieve existing product and product tag IDs
	var salesGroups []models.SalesGroup
	var paymentMethods []models.PaymentMethod
	var verifiers []models.Verifier
	var buyerTypes []models.BuyerType
	db.Find(&paymentMethods)
	db.Find(&salesGroups)
	db.Find(&verifiers)
	db.Find(&buyerTypes)
	for _, paymentMethod := range paymentMethods {
		for _, salesGroup := range salesGroups {
			db.Model(&salesGroup).Association("PaymentMethodsSet").Append(&paymentMethod)
		}
	}
	for _, verifier := range verifiers {
		for _, salesGroup := range salesGroups {
			db.Model(&salesGroup).Association("VerifierSet").Append(&verifier)
		}
	}
	for _, buyerType := range buyerTypes {
		for _, salesGroup := range salesGroups {
			db.Model(&salesGroup).Association("BuyerTypesSet").Append(&buyerType)
		}
	}
}
