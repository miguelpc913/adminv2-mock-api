package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedVenues(db *gorm.DB) {
	venues := []models.Venue{
		{
			VenueId: 1,
			Name:    "Conference Hall A",
		},
		{
			VenueId: 2,
			Name:    "Main Auditorium",
		},
		{
			VenueId: 3,
			Name:    "Outdoor Amphitheater",
		},
		{
			VenueId: 4,
			Name:    "Banquet Room B",
		},
		{
			VenueId: 5,
			Name:    "Exhibition Center",
		},
	}

	result := db.Create(&venues)
	if result.Error != nil {
		fmt.Println("Error occurred while seeding venues:", result.Error)
	}
}
