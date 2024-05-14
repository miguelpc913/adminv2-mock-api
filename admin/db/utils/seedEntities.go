package dbHelpers

import (
	"errors"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"github.com/tiqueteo/adminv2-mock-api/db/seeds"
	"gorm.io/gorm"
)

func isEntityEmpty(db *gorm.DB, entity interface{}) bool {
	err := db.First(&entity).Error
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func isDBEmpty(db *gorm.DB) bool {
	entities := []interface{}{
		&models.BuyerType{},
		&models.Verifier{},
		&models.PaymentMethod{},
		&models.Product{},
		&models.ProductTag{},
		&models.SalesGroupHtml{},
		&models.SalesGroup{},
		&models.Venue{},
		&models.ProductInfoType{},
		&models.ProductInfo{},
		&models.RecommendationRule{},
		&models.Promotion{},
		&models.AffiliateItem{},
		&models.AffiliateAgreement{},
		&models.Affiliate{},
	}

	for _, entity := range entities {
		if !isEntityEmpty(db, entity) {
			return false
		}
	}
	return true
}

func SeedEntities(db *gorm.DB) {
	if isDBEmpty(db) {
		seeds.SeedBuyerType(db)
		seeds.SeedVerifier(db)
		seeds.SeedPaymentMethod(db)
		seeds.SeedProduct(db)
		seeds.SeedProductTag(db)
		seeds.SeedSalesGroupHtml(db)
		seeds.SeedVenues(db)
		seeds.SeedSalesGroups(db)
		seeds.SeedProductInfoTypes(db)
		seeds.SeedProductInfos(db)
		seeds.SeedRecommendationRules(db)
		seeds.SeedPromotions(db)
		seeds.SeedPromotionPrices(db)
		seeds.SeedPromotionCodes(db)
		seeds.SeedAffiliateItems(db)
		seeds.SeedAffiliateAgreements(db)
		seeds.SeedAffiliates(db)
		// Seed associations
		seeds.SeedSalesGroupsAssociations(db)
		seeds.SeedProductTagAssociation(db)
		seeds.SeedProductInfoAssociations(db)
		seeds.SeedRecommendationRulesAssociations(db)
		seeds.SeedPromotionPrices(db)
		seeds.SeedPromotionCodes(db)
		seeds.SeedPromotionAssociations(db)
		seeds.SeedAffiliateItemsAssociations(db)
		seeds.SeedAffiliateAgreementAssociations(db)
		seeds.SeedAffiliateAssociations(db)
	}
}
