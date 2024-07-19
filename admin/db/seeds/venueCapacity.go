package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedVenueCapacities(db *gorm.DB) {
	venues := []models.VenueCapacity{
		{
			Name: `{ "es": "Visita Libre", "en": "Self Guided Tour" }`,
		},
		{
			Name: `{ "es": "Guiada Español", "en": "Guided Spanish" }`,
		},
		{
			Name: "Riverside Convention Hall",
		},
		{
			Name: `{ "es": "Guiada Inglés", "en": "English Guided" }`,
		},
		{
			Name: `{ "es": "Open Libre", "en": "Open" }`,
		},
	}

	result := db.Create(&venues)
	if result.Error != nil {
		fmt.Println("Error occurred while seeding venues:", result.Error)
	}
}
