package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedExtra(db *gorm.DB) {
	taxes := []models.Tax{
		{
			TaxId: 1,
			Name:  "VAT",
		},
		{
			TaxId: 2,
			Name:  "Sales Tax",
		},
		{
			TaxId: 3,
			Name:  "Service Tax",
		},
	}
	result := db.Create(&taxes)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of pricings: taxes")
	}
	extras := []models.Extra{
		{
			ExtraId:          1,
			Name:             "Audioguia",
			Status:           true,
			TaxId:            1,
			HideFromSummary:  false,
			ManagesStock:     true,
			ManagesStockType: "optional",
			Stock:            100,
		},
		{
			ExtraId:          2,
			Name:             "Visita guiada",
			Status:           true,
			TaxId:            2,
			HideFromSummary:  true,
			ManagesStock:     false,
			ManagesStockType: "optional",
			Stock:            0,
		},
		{
			ExtraId:          3,
			Name:             "Visita libre",
			Status:           false,
			TaxId:            3,
			HideFromSummary:  false,
			ManagesStock:     false,
			ManagesStockType: "optional",
			Stock:            0,
		},
		{
			ExtraId:          4,
			Name:             "Mediador/a",
			Status:           true,
			TaxId:            1,
			HideFromSummary:  true,
			ManagesStock:     false,
			ManagesStockType: "mandatory",
			Stock:            10,
		},
		{
			ExtraId:          5,
			Name:             "Con copa de vino",
			Status:           true,
			TaxId:            2,
			HideFromSummary:  false,
			ManagesStock:     true,
			ManagesStockType: "mandatory",
			Stock:            50,
		},
	}
	result = db.Create(&extras)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of pricings: extras")
	}
}
