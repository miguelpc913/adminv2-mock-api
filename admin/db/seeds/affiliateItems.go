package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedAffiliateItems(db *gorm.DB) {
	affiliateItems := []models.AffiliateItem{
		{
			Status:        true,
			ItemName:      "Item 1",
			ItemType:      "custom_url",
			ItemLinkedUrl: `{"en": "http://example.com/item1_en", "es": "http://example.com/item1_es", "fr": "http://example.com/item1_fr"}`,
			ItemResource:  nil,
			IsGeneric:     false,
		},
		{
			Status:        true,
			ItemName:      "Item 2",
			ItemType:      "custom_url",
			ItemLinkedUrl: `{"en": "http://example.com/item2_en", "es": "http://example.com/item2_es", "fr": "http://example.com/item2_fr"}`,
			ItemResource:  nil,
			IsGeneric:     false,
		},
		{
			Status:        true,
			ItemName:      "Banner 1",
			ItemType:      "banner",
			ItemLinkedUrl: `{"en": "http://example.com/banner1", "es": "http://example.com/banner1", "fr": "http://example.com/banner1"}`,
			ItemResource:  stringPtr(`{"en": "banner1.jpg", "es": "banner1.jpg", "fr": "banner1.jpg"}`),
			IsGeneric:     false,
		},
		{
			Status:        true,
			ItemName:      "Item 3",
			ItemType:      "custom_url",
			ItemLinkedUrl: `{"en": "http://example.com/item3_en", "es": "http://example.com/item3_es", "fr": "http://example.com/item3_fr"}`,
			ItemResource:  nil,
			IsGeneric:     false,
		},
		{
			Status:        true,
			ItemName:      "Banner 2",
			ItemType:      "banner",
			ItemLinkedUrl: `{"en": "http://example.com/banner1", "es": "http://example.com/banner1", "fr": "http://example.com/banner1"}`,
			ItemResource:  stringPtr(`{"en": "banner2.jpg", "es": "banner2.jpg", "fr": "banner2.jpg"}`),
			IsGeneric:     false,
		},
	}

	result := db.Create(&affiliateItems)
	if result.Error != nil {
		fmt.Println("Error seeding affiliate_items:", result.Error)
	}
}
