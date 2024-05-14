package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedSalesGroupHtml(db *gorm.DB) {
	htmls := []models.SalesGroupHtml{
		{
			Status:       true,
			Name:         `{ "en": "HTML Content A", "es": "Contenido HTML A", "ca": "Contingut HTML A" }`,
			Icon:         "icon-a",
			Slug:         `{ "en": "html-content-a", "es": "contenido-html-a", "ca": "contingut-html-a" }`,
			InnerHTML:    `{ "en": "HTML A Details", "es": "Detalles HTML A", "ca": "Detalls HTML A" }`,
			DisplayOrder: 1,
		},
		{
			Status:       true,
			Name:         `{ "en": "HTML Content B", "es": "Contenido HTML B", "ca": "Contingut HTML B" }`,
			Icon:         "icon-b",
			Slug:         `{ "en": "html-content-b", "es": "contenido-html-b", "ca": "contingut-html-b" }`,
			InnerHTML:    `{ "en": "HTML B Details", "es": "Detalles HTML B", "ca": "Detalls HTML B" }`,
			DisplayOrder: 2,
		},
		{
			Status:       true,
			Name:         `{ "en": "HTML Content C", "es": "Contenido HTML C", "ca": "Contingut HTML C" }`,
			Icon:         "icon-c",
			Slug:         `{ "en": "html-content-c", "es": "contenido-html-c", "ca": "contingut-html-c" }`,
			InnerHTML:    `{ "en": "HTML C Details", "es": "Detalles HTML C", "ca": "Detalls HTML C" }`,
			DisplayOrder: 3,
		},
		{
			Status:       true,
			Name:         `{ "en": "HTML Content D", "es": "Contenido HTML D", "ca": "Contingut HTML D" }`,
			Icon:         "icon-d",
			Slug:         `{ "en": "html-content-d", "es": "contenido-html-d", "ca": "contingut-html-d" }`,
			InnerHTML:    `{ "en": "HTML D Details", "es": "Detalles HTML D", "ca": "Detalls HTML D" }`,
			DisplayOrder: 4,
		},
		{
			Status:       true,
			Name:         `{ "en": "HTML Content E", "es": "Contenido HTML E", "ca": "Contingut HTML E" }`,
			Icon:         "icon-e",
			Slug:         `{ "en": "html-content-e", "es": "contenido-html-e", "ca": "contingut-html-e" }`,
			InnerHTML:    `{ "en": "HTML E Details", "es": "Detalles HTML E", "ca": "Detalls HTML E" }`,
			DisplayOrder: 5,
		},
	}

	result := db.Create(htmls)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of sales group HTMLs")
	}
}
