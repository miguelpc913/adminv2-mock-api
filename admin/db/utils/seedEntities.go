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

func SeedEntities(db *gorm.DB) {
	entityEmptyMap := map[string]bool{}
	if isEntityEmpty(db, &models.BuyerType{}) {
		seeds.SeedBuyerType(db)
		entityEmptyMap["BuyerType"] = true
	}
	if isEntityEmpty(db, &models.Verifier{}) {
		seeds.SeedVerifier(db)
		entityEmptyMap["Verifier"] = true
	}
	if isEntityEmpty(db, &models.PaymentMethod{}) {
		seeds.SeedPaymentMethod(db)
		entityEmptyMap["PaymentMethod"] = true
	}
	if isEntityEmpty(db, &models.Product{}) {
		seeds.SeedProduct(db)
		entityEmptyMap["Product"] = true
	}
	if isEntityEmpty(db, &models.ProductTag{}) {
		seeds.SeedProductTag(db)
		entityEmptyMap["ProductTag"] = true
	}
	if isEntityEmpty(db, &models.SalesGroupHtml{}) {
		seeds.SeedSalesGroupHtml(db)
		entityEmptyMap["SalesGroupHtml"] = true
	}
	if isEntityEmpty(db, &models.VenueCapacity{}) {
		seeds.SeedVenueCapacities(db)
		entityEmptyMap["VenueCapacity"] = true
	}
	if isEntityEmpty(db, &models.Venue{}) {
		seeds.SeedVenues(db)
		entityEmptyMap["Venues"] = true
	}
	if isEntityEmpty(db, &models.SalesGroup{}) {
		seeds.SeedSalesGroups(db)
		entityEmptyMap["SalesGroup"] = true
	}
	if isEntityEmpty(db, &models.ProductInfoType{}) {
		seeds.SeedProductInfoTypes(db)
		entityEmptyMap["ProductInfoType"] = true
	}
	if isEntityEmpty(db, &models.ProductInfo{}) {
		seeds.SeedProductInfos(db)
		entityEmptyMap["ProductInfo"] = true
	}
	if isEntityEmpty(db, &models.RecommendationRule{}) {
		seeds.SeedRecommendationRules(db)
		entityEmptyMap["RecommendationRule"] = true
	}
	if isEntityEmpty(db, &models.Promotion{}) {
		seeds.SeedPromotions(db)
		entityEmptyMap["Promotion"] = true
	}
	if isEntityEmpty(db, &models.AffiliateItem{}) {
		seeds.SeedAffiliateItems(db)
		entityEmptyMap["AffiliateItem"] = true
	}
	if isEntityEmpty(db, &models.AffiliateAgreement{}) {
		seeds.SeedAffiliateAgreements(db)
		entityEmptyMap["AffiliateAgreement"] = true
	}
	if isEntityEmpty(db, &models.Affiliate{}) {
		seeds.SeedAffiliates(db)
		entityEmptyMap["Affiliate"] = true
	}
	if isEntityEmpty(db, &models.BoxOffice{}) {
		seeds.SeedBoxOffice(db)
		entityEmptyMap["BoxOffice"] = true
	}

	if (isEntityEmpty(db, &models.MainPricing{}) && isEntityEmpty(db, &models.SpecificPricing{}) && isEntityEmpty(db, &models.ProductVenueBuyerTypes{}) && isEntityEmpty(db, &models.ProductExtraBuyerTypes{})) {
		seeds.SeedBoxOffice(db)
		entityEmptyMap["Pricings"] = true
	}

	// Seed associations
	if entityEmptyMap["SalesGroup"] {
		seeds.SeedSalesGroupsAssociations(db)
	}
	if entityEmptyMap["ProductTag"] {
		seeds.SeedProductTagAssociation(db)
	}
	if entityEmptyMap["ProductInfo"] {
		seeds.SeedProductInfoAssociations(db)
	}
	if entityEmptyMap["RecommendationRule"] {
		seeds.SeedRecommendationRulesAssociations(db)
	}
	if entityEmptyMap["Promotion"] {
		seeds.SeedPromotionPrices(db)
		seeds.SeedPromotionCodes(db)
		seeds.SeedPromotionAssociations(db)
	}
	if entityEmptyMap["AffiliateItem"] {
		seeds.SeedAffiliateItemsAssociations(db)
	}
	if entityEmptyMap["AffiliateAgreement"] {
		seeds.SeedAffiliateAgreementAssociations(db)
	}
	if entityEmptyMap["Affiliate"] {
		seeds.SeedAffiliateAssociations(db)
	}
	if entityEmptyMap["BoxOffice"] {
		seeds.SeedBoxOfficeAssociations(db)
		seeds.SeedBoxOfficeLanguages(db)
	}
	if entityEmptyMap["Pricings"] {
		seeds.SeedMainPricing(db)
	}

}
