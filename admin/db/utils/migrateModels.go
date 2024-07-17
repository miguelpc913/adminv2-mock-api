package dbHelpers

import (
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.BuyerType{})
	db.AutoMigrate(&models.SalesGroup{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductTag{})
	db.AutoMigrate(&models.SalesGroupHtml{})
	db.AutoMigrate(&models.Verifier{})
	db.AutoMigrate(&models.PaymentMethod{})
	db.AutoMigrate(&models.Venue{})
	db.AutoMigrate(&models.ProductInfo{})
	db.AutoMigrate(&models.ProductInfoType{})
	db.AutoMigrate(&models.RecommendationRule{})
	db.AutoMigrate(&models.Promotion{})
	db.AutoMigrate(&models.PromotionPrice{})
	db.AutoMigrate(&models.PromotionalCode{})
	db.AutoMigrate(&models.AffiliateItem{})
	db.AutoMigrate(&models.AffiliateAgreement{})
	db.AutoMigrate(&models.Affiliate{})
	db.AutoMigrate(&models.BoxOffice{})
	db.AutoMigrate(&models.AllowedAppLanguages{})
	db.AutoMigrate(&models.AllowedTicketLanguages{})
	db.AutoMigrate(&models.ProductExtraBuyerTypes{})
	db.AutoMigrate(&models.ProductVenueBuyerTypes{})
	db.AutoMigrate(&models.SpecificPricing{})
	db.AutoMigrate(&models.MainPricing{})
}
