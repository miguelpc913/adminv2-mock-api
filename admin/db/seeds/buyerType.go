package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedBuyerType(db *gorm.DB) {
	buyerTypes := []models.BuyerType{
		{
			Name:                  `{"en": "Standard", "es": "Estándar", "ca": "Estàndard"}`,
			ShortName:             "STD",
			Description:           `{"en": "Standard buyer type", "es": "Tipo de comprador estándar", "ca": "Tipus de comprador estàndard"}`,
			Status:                true,
			DiscountRate:          0.10,
			DiscountAmount:        5.00,
			NoHandleAvailability:  false,
			Pax:                   100,
			RequiredPdaColorAlert: false,
			AlertColor:            "none",
			RequiredPdaSoundAlert: false,
			AlertSound:            "none",
		},
		{
			Name:                  `{"en": "Premium", "es": "Premium", "ca": "Premium"}`,
			ShortName:             "PRM",
			Description:           `{"en": "Premium buyer type", "es": "Tipo de comprador premium", "ca": "Tipus de comprador premium"}`,
			Status:                true,
			DiscountRate:          0.20,
			DiscountAmount:        10.00,
			NoHandleAvailability:  true,
			Pax:                   200,
			RequiredPdaColorAlert: true,
			AlertColor:            "red",
			RequiredPdaSoundAlert: true,
			AlertSound:            "beep",
		},
		{
			Name:                  `{"en": "Wholesale", "es": "Mayorista", "ca": "Majorista"}`,
			ShortName:             "WHL",
			Description:           `{"en": "Buyer for wholesale", "es": "Comprador para mayoristas", "ca": "Comprador per a majoristes"}`,
			Status:                false,
			DiscountRate:          0.15,
			DiscountAmount:        15.00,
			NoHandleAvailability:  true,
			Pax:                   300,
			RequiredPdaColorAlert: false,
			AlertColor:            "none",
			RequiredPdaSoundAlert: false,
			AlertSound:            "none",
		},
		{
			Name:                  `{"en": "Retail", "es": "Minorista", "ca": "Minorista"}`,
			ShortName:             "RTL",
			Description:           `{"en": "Buyer for retail", "es": "Comprador para minoristas", "ca": "Comprador per a minoristes"}`,
			Status:                true,
			DiscountRate:          0.05,
			DiscountAmount:        2.50,
			NoHandleAvailability:  false,
			Pax:                   50,
			RequiredPdaColorAlert: true,
			AlertColor:            "blue",
			RequiredPdaSoundAlert: true,
			AlertSound:            "ring",
		},
		{
			Name:                  `{"en": "Occasional", "es": "Ocasional", "ca": "Occasional"}`,
			ShortName:             "OCC",
			Description:           `{"en": "Occasional buyer type", "es": "Tipo de comprador ocasional", "ca": "Tipus de comprador occasional"}`,
			Status:                false,
			DiscountRate:          0.08,
			DiscountAmount:        3.00,
			NoHandleAvailability:  false,
			Pax:                   75,
			RequiredPdaColorAlert: false,
			AlertColor:            "none",
			RequiredPdaSoundAlert: false,
			AlertSound:            "none",
		},
	}
	result := db.Create(buyerTypes)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of buyer type")
	}
}
