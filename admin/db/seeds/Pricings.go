package seeds

import (
	"fmt"
	"time"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedMainPricing(db *gorm.DB) {
	mainPricings := []models.MainPricing{
		{
			MainPricingId: 1,
			StartDatetime: "2024-01-01T00:00:00Z",
			EndDateTime:   "2024-12-31T23:59:59Z",
			ProductId:     2,
		},
		{
			MainPricingId: 2,
			StartDatetime: "2025-01-01T00:00:00Z",
			EndDateTime:   "2025-12-31T23:59:59Z",
			ProductId:     2,
		},
	}

	result := db.Create(&mainPricings)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of mainPricings")
	}

	specificPricings := []models.SpecificPricing{
		{
			PricingId:     1,
			MainPricingId: 1,
			Name:          "Tarifario Base",
			Priority:      1,
			Weekdays:      []int{1, 2, 3, 4, 5, 6, 7},
			EnabledDates:  []string{},
			Default:       true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			PricingId:     2,
			MainPricingId: 1,
			Name:          "Tarifario Especial 2024 - Fin de Semana",
			Priority:      2,
			Weekdays:      []int{6, 7},
			EnabledDates:  []string{},
			Default:       false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			PricingId:     3,
			MainPricingId: 2,
			Name:          "Tarifario Base",
			Priority:      1,
			Weekdays:      []int{1, 2, 3, 4, 5, 6, 7},
			EnabledDates:  []string{},
			Default:       true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			PricingId:     4,
			MainPricingId: 2,
			Name:          "Tarifario Especial 2025 - Temporada Alta",
			Priority:      2,
			Weekdays:      []int{1, 2, 3, 4, 5, 6, 7},
			EnabledDates:  []string{"2025-06-01", "2025-06-02", "2025-06-03"},
			Default:       false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	result = db.Create(&specificPricings)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of specific pricings")
	}

	productVenueBuyerTypes := []models.ProductVenueBuyerTypes{
		{
			ProductVenueBuyerTypeId: 1,
			ProductId:               1,
			PricingId:               1,
			VenueId:                 1,
			BuyerTypeId:             1,
			Price:                   100.0,
			HasDiscount:             false,
		},
		{
			ProductVenueBuyerTypeId: 2,
			ProductId:               1,
			PricingId:               1,
			VenueId:                 2,
			BuyerTypeId:             2,
			Price:                   150.0,
			HasDiscount:             true,
		},
		{
			ProductVenueBuyerTypeId: 3,
			ProductId:               1,
			PricingId:               2,
			VenueId:                 3,
			BuyerTypeId:             3,
			Price:                   200.0,
			HasDiscount:             false,
		},
		{
			ProductVenueBuyerTypeId: 4,
			ProductId:               1,
			PricingId:               2,
			VenueId:                 4,
			BuyerTypeId:             4,
			Price:                   250.0,
			HasDiscount:             true,
		},
		{
			ProductVenueBuyerTypeId: 5,
			ProductId:               2,
			PricingId:               3,
			VenueId:                 1,
			BuyerTypeId:             1,
			Price:                   220.0,
			HasDiscount:             false,
		},
		{
			ProductVenueBuyerTypeId: 6,
			ProductId:               2,
			PricingId:               3,
			VenueId:                 2,
			BuyerTypeId:             2,
			Price:                   230.0,
			HasDiscount:             true,
		},
		{
			ProductVenueBuyerTypeId: 7,
			ProductId:               2,
			PricingId:               4,
			VenueId:                 3,
			BuyerTypeId:             3,
			Price:                   240.0,
			HasDiscount:             false,
		},
		{
			ProductVenueBuyerTypeId: 8,
			ProductId:               2,
			PricingId:               4,
			VenueId:                 4,
			BuyerTypeId:             4,
			Price:                   250.0,
			HasDiscount:             true,
		},
	}
	result = db.Create(&productVenueBuyerTypes)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of pricings: product/venue/buyertypes")
	}

	productExtraBuyerTypes := []models.ProductExtraBuyerTypes{
		{
			ProductExtraBuyerTypeId: 1,
			ProductId:               1,
			ExtraId:                 1,
			PricingId:               1,
			BuyerTypeId:             1,
			Price:                   50.0,
			HasDiscount:             false,
		},
		{
			ProductExtraBuyerTypeId: 2,
			ProductId:               1,
			ExtraId:                 2,
			PricingId:               1,
			BuyerTypeId:             2,
			Price:                   75.0,
			HasDiscount:             true,
		},
		{
			ProductExtraBuyerTypeId: 3,
			ProductId:               1,
			ExtraId:                 3,
			PricingId:               2,
			BuyerTypeId:             3,
			Price:                   80.0,
			HasDiscount:             false,
		},
		{
			ProductExtraBuyerTypeId: 4,
			ProductId:               1,
			ExtraId:                 4,
			PricingId:               2,
			BuyerTypeId:             4,
			Price:                   90.0,
			HasDiscount:             true,
		},
		{
			ProductExtraBuyerTypeId: 5,
			ProductId:               2,
			ExtraId:                 1,
			PricingId:               3,
			BuyerTypeId:             1,
			Price:                   60.0,
			HasDiscount:             false,
		},
		{
			ProductExtraBuyerTypeId: 6,
			ProductId:               2,
			ExtraId:                 2,
			PricingId:               3,
			BuyerTypeId:             2,
			Price:                   70.0,
			HasDiscount:             true,
		},
		{
			ProductExtraBuyerTypeId: 7,
			ProductId:               2,
			ExtraId:                 3,
			PricingId:               4,
			BuyerTypeId:             3,
			Price:                   80.0,
			HasDiscount:             false,
		},
		{
			ProductExtraBuyerTypeId: 8,
			ProductId:               2,
			ExtraId:                 4,
			PricingId:               4,
			BuyerTypeId:             4,
			Price:                   90.0,
			HasDiscount:             true,
		},
	}
	result = db.Create(&productExtraBuyerTypes)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of pricings: product/extra/buyertypes")
	}
}
