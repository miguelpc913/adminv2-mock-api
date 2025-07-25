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
			StartDate: "2024-01-01",
			EndDate:   "2024-12-31",
			Color:     "#93CEF8",
			Name:      "Tarifario base 1",
			ProductId: 2,
		},
		{
			StartDate: "2025-01-01",
			EndDate:   "2025-12-31",
			Color:     "#98DC98",
			Name:      "Tarifario base 2",
			ProductId: 2,
		},
	}

	result := db.Create(&mainPricings)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of mainPricings")
	}

	specificPricings := []models.SpecificPricing{
		{

			MainPricingId: 1,
			Name:          "Tarifario Base",
			Priority:      99999,
			Weekdays:      []int{1, 2, 3, 4, 5, 6, 7},
			EnabledDates:  []string{},
			StartHour:     stringPtr("08:00:00"),
			EndHour:       stringPtr("20:00:00"),
			Default:       true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{

			MainPricingId: 1,
			Name:          "Tarifario Especial 2024 - Fin de Semana",
			Priority:      1,
			Weekdays:      []int{6, 7},
			EnabledDates:  []string{},
			StartHour:     stringPtr("08:00:00"),
			EndHour:       stringPtr("20:00:00"),
			Default:       false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},

		{

			MainPricingId: 2,
			Name:          "Tarifario Base",
			Priority:      99999,
			Weekdays:      []int{1, 2, 3, 4, 5, 6, 7},
			EnabledDates:  []string{},
			StartHour:     stringPtr("08:00:00"),
			EndHour:       stringPtr("20:00:00"),

			Default:   true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{

			MainPricingId: 2,
			Name:          "Tarifario Especial 2025 - Temporada Alta",
			Priority:      1,
			Weekdays:      []int{1, 2, 3, 4, 5, 6, 7},
			EnabledDates:  []string{"2025-06-01", "2025-06-02", "2025-06-03"},
			StartHour:     stringPtr("08:00:00"),
			EndHour:       stringPtr("20:00:00"),
			Default:       false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	result = db.Create(&specificPricings)
	if result.Error != nil {
		fmt.Print(result.Error.Error())
	}
	// RecurrentTimes
	recurrentTimes := []models.RecurrentTime{
		{

			PricingId: 1,
			Minutes:   []int{0, 30},
			Hours:     []int{10, 13, 18},
		},
		{

			PricingId: 2,
			Minutes:   []int{15, 45},
			Hours:     []int{9, 12, 15},
		},
	}
	result = db.Create(&recurrentTimes)
	if result.Error != nil {
		fmt.Print(result.Error.Error())
	}

	// DynamicPricingConfigurations
	dynamicPricingConfigs := []models.DynamicPricingConfiguration{
		{

			PricingId: 3,
			Type:      "event_range",
			StartHour: stringPtr("10:00:00"),
			EndHour:   stringPtr("18:00:00"),
		},
		{

			PricingId: 4,
			Type:      "event_range",
			StartHour: stringPtr("12:00:00"),
			EndHour:   stringPtr("20:00:00"),
		},
	}
	result = db.Create(&dynamicPricingConfigs)
	if result.Error != nil {
		fmt.Print(result.Error.Error())
	}

	// OccupancyRanges
	occupancyRanges := []models.OccupancyRange{
		{DynamicPricingConfigurationId: 1, Start: 0, End: 25},
		{DynamicPricingConfigurationId: 1, Start: 26, End: 75},
		{DynamicPricingConfigurationId: 1, Start: 76, End: 100},
		{DynamicPricingConfigurationId: 2, Start: 0, End: 50},
		{DynamicPricingConfigurationId: 2, Start: 51, End: 100},
	}
	result = db.Create(&occupancyRanges)
	if result.Error != nil {
		fmt.Print(result.Error.Error())
	}

	productVenueBuyerTypes := []models.ProductVenueBuyerTypes{
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 1,
			PricingId:   1,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 1,
			PricingId:   1,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 2,
			PricingId:   1,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 2,
			PricingId:   1,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 3,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 3,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 4,
			PricingId:   1,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 4,
			PricingId:   1,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 5,
			PricingId:   1,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 5,
			PricingId:   1,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 6,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 6,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 7,
			PricingId:   1,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 7,
			PricingId:   1,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 1,
			PricingId:   2,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 1,
			PricingId:   2,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 2,
			PricingId:   2,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 2,
			PricingId:   2,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 3,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 3,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 4,
			PricingId:   2,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 4,
			PricingId:   2,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 5,
			PricingId:   2,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 5,
			PricingId:   2,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 6,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 6,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 7,
			PricingId:   2,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 7,
			PricingId:   2,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 1,
			PricingId:   3,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 1,
			PricingId:   3,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 2,
			PricingId:   3,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 2,
			PricingId:   3,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 3,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 3,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 4,
			PricingId:   3,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 4,
			PricingId:   3,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 5,
			PricingId:   3,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 5,
			PricingId:   3,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 6,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 6,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 7,
			PricingId:   3,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 7,
			PricingId:   3,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 1,
			PricingId:   4,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 1,
			PricingId:   4,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 2,
			PricingId:   4,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 2,
			PricingId:   4,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 3,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 3,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 4,
			PricingId:   4,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 4,
			PricingId:   4,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 5,
			PricingId:   4,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 5,
			PricingId:   4,
			HasDiscount: false,
			Price:       10,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 6,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 6,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			VenueId:     1,
			BuyerTypeId: 7,
			PricingId:   4,
			HasDiscount: false,
			Price:       4,
		},
		{
			ProductId:   2,
			VenueId:     2,
			BuyerTypeId: 7,
			PricingId:   4,
			HasDiscount: false,
			Price:       10,
		},
	}
	result = db.Create(&productVenueBuyerTypes)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of pricings: product/venue/buyertypes")
	}

	productExtraBuyerTypes := []models.ProductExtraBuyerTypes{
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 1,
			PricingId:   1,
			HasDiscount: false,
			Price:       26,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 2,
			PricingId:   1,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 3,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 4,
			PricingId:   1,
			HasDiscount: false,
			Price:       18,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 5,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 6,
			PricingId:   1,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 7,
			PricingId:   1,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 1,
			PricingId:   2,
			HasDiscount: false,
			Price:       26,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 2,
			PricingId:   2,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 3,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 4,
			PricingId:   2,
			HasDiscount: false,
			Price:       18,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 5,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 6,
			PricingId:   2,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 7,
			PricingId:   2,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 1,
			PricingId:   3,
			HasDiscount: false,
			Price:       26,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 2,
			PricingId:   3,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 3,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 4,
			PricingId:   3,
			HasDiscount: false,
			Price:       18,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 5,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 6,
			PricingId:   3,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 7,
			PricingId:   3,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 1,
			PricingId:   4,
			HasDiscount: false,
			Price:       26,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 2,
			PricingId:   4,
			HasDiscount: false,
			Price:       24,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 3,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 4,
			PricingId:   4,
			HasDiscount: false,
			Price:       18,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 5,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 6,
			PricingId:   4,
			HasDiscount: false,
			Price:       0,
		},
		{
			ProductId:   2,
			ExtraId:     1,
			BuyerTypeId: 7,
			PricingId:   4,
			HasDiscount: false,
			Price:       24,
		},
	}
	result = db.Create(&productExtraBuyerTypes)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of pricings: product/extra/buyertypes")
	}
}
