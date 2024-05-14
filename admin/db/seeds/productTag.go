package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedProductTag(db *gorm.DB) {
	productTags := []models.ProductTag{
		{
			Name:                   `{ "en": "New Arrival", "es": "Recién Llegado", "ca": "Nouvingut" }`,
			Status:                 true,
			ProductCategoryGeneral: false,
		},
		{
			Name:                   `{ "en": "Bestseller", "es": "Más Vendidos", "ca": "Més Venuts" }`,
			Status:                 true,
			ProductCategoryGeneral: false,
		},
		{
			Name:                   `{ "en": "Discounted", "es": "Con Descuento", "ca": "Amb Descompte" }`,
			Status:                 true,
			ProductCategoryGeneral: false,
		},
		{
			Name:                   `{ "en": "Limited Edition", "es": "Edición Limitada", "ca": "Edició Limitada" }`,
			Status:                 true,
			ProductCategoryGeneral: false,
		},
		{
			Name:                   `{ "en": "Online Special", "es": "Especial Online", "ca": "Especial Online" }`,
			Status:                 true,
			ProductCategoryGeneral: false,
		},
	}

	result := db.Create(productTags)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of product tags")
	}
}
