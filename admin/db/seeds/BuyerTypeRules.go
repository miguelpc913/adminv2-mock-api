package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	modelHelpers "github.com/tiqueteo/adminv2-mock-api/db/models/helpers"
	"gorm.io/gorm"
)

func SeedBuyerTypeRules(db *gorm.DB) {
	// Shared Vars template
	vars := models.Vars{
		X: modelHelpers.JSONIntSLice{1, 2, 3},
		Y: modelHelpers.JSONIntSLice{3, 4},
		M: ptrInt(5),
		N: ptrInt(6),
	}
	fmt.Println("injected buyer type rules")
	buyerRuleRules := []models.BuyerTypeRule{
		{
			Status:                  true,
			Name:                    "Standard",
			BuyerTypeRuleTemplateID: 1,

			ErrorMessage: `{"en": "Standard error message", "es": "Mensaje de error estándar", "ca": "Missatge d'error estàndard"}`,
			Priority:     1,
		},
		{
			Status:                  true,
			Name:                    "Premium",
			BuyerTypeRuleTemplateID: 2,

			ErrorMessage: `{"en": "Premium error message", "es": "Mensaje de error premium", "ca": "Missatge d'error premium"}`,
			Priority:     2,
		},
		{
			Status:                  false,
			Name:                    "Wholesale",
			BuyerTypeRuleTemplateID: 3,

			ErrorMessage: `{"en": "Wholesale error", "es": "Error de mayorista", "ca": "Error majorista"}`,
			Priority:     3,
		},
		{
			Status:                  true,
			Name:                    "Retail",
			BuyerTypeRuleTemplateID: 4,

			ErrorMessage: `{"en": "Retail rule error", "es": "Error de minorista", "ca": "Error de detallista"}`,
			Priority:     4,
		},
		{
			Status:                  false,
			Name:                    "Occasional",
			BuyerTypeRuleTemplateID: 5,

			ErrorMessage: `{"en": "Occasional error", "es": "Error ocasional", "ca": "Error esporàdic"}`,
			Priority:     5,
		},
		{
			Status:                  true,
			Name:                    "Corporate",
			BuyerTypeRuleTemplateID: 5,

			ErrorMessage: `{"en": "Corporate error message", "es": "Mensaje de error corporativo", "ca": "Missatge d'error corporatiu"}`,
			Priority:     6,
		},
		{
			Status:                  true,
			Name:                    "VIP",
			BuyerTypeRuleTemplateID: 5,

			ErrorMessage: `{"en": "VIP rule error", "es": "Error de VIP", "ca": "Error VIP"}`,
			Priority:     7,
		},
	}

	result := db.Create(&buyerRuleRules)
	if result.Error != nil {
		fmt.Println("Error seeding buyer_type_rules:", result.Error)
	}

	var buyerTypeRules []models.BuyerTypeRule
	db.Find(&buyerTypeRules)
	var products []models.Product
	db.Find(&products)
	for _, product := range products {
		for _, buyerTypeRule := range buyerTypeRules {
			db.Model(&buyerTypeRule).Association("ProductSet").Append(&product)
		}
	}

	for _, buyerTypeRule := range buyerTypeRules {
		varsCopy := vars // create a copy for each rule to avoid reuse
		varsCopy.BuyerTypeRuleID = buyerTypeRule.BuyerTypeRuleID
		if err := db.Create(&varsCopy).Error; err != nil {
			fmt.Printf("Error saving Vars for rule %s: %v\n", buyerTypeRule.Name, err)
		}
	}

}

func ptrInt(i int) *int {
	return &i
}
