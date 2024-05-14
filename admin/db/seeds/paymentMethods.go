package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedPaymentMethod(db *gorm.DB) {
	paymentMethods := []models.PaymentMethod{
		{
			Name: `{ "en": "Credit Card", "es": "Tarjeta de Crédito", "ca": "Targeta de Crèdit" }`,
		},
		{
			Name: `{ "en": "PayPal", "es": "PayPal", "ca": "PayPal" }`,
		},
		{
			Name: `{ "en": "Bank Transfer", "es": "Transferencia Bancaria", "ca": "Transferència Bancària" }`,
		},
		{
			Name: `{ "en": "Cash", "es": "Efectivo", "ca": "Efectiu" }`,
		},
		{
			Name: `{ "en": "Cryptocurrency", "es": "Criptomoneda", "ca": "Criptomoneda" }`,
		},
	}

	result := db.Create(paymentMethods)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of payment methods")
	}
}
