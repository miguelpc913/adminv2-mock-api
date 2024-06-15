package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedBoxOfficeLanguages(db *gorm.DB) {
	allowedTicketLanguages := []models.AllowedTicketLanguages{
		{BoxOfficeId: 1, LanguageCode: "en", DisplayOrder: 1},
		{BoxOfficeId: 1, LanguageCode: "es", DisplayOrder: 2},
		{BoxOfficeId: 1, LanguageCode: "it", DisplayOrder: 3},
		{BoxOfficeId: 2, LanguageCode: "es", DisplayOrder: 1},
		{BoxOfficeId: 3, LanguageCode: "fr", DisplayOrder: 1},
		{BoxOfficeId: 4, LanguageCode: "mx", DisplayOrder: 1},
		{BoxOfficeId: 5, LanguageCode: "it", DisplayOrder: 1},
	}
	allowedAppLanguages := []models.AllowedAppLanguages{
		{BoxOfficeId: 1, LanguageCode: "en", DisplayOrder: 1},
		{BoxOfficeId: 1, LanguageCode: "es", DisplayOrder: 2},
		{BoxOfficeId: 1, LanguageCode: "it", DisplayOrder: 3},
		{BoxOfficeId: 2, LanguageCode: "es", DisplayOrder: 1},
		{BoxOfficeId: 3, LanguageCode: "fr", DisplayOrder: 1},
		{BoxOfficeId: 4, LanguageCode: "mx", DisplayOrder: 1},
		{BoxOfficeId: 5, LanguageCode: "it", DisplayOrder: 1},
	}
	result := db.Create(allowedTicketLanguages)
	if result.Error != nil {
		fmt.Println("Error seeding allowed ticket languages:", result.Error)
	}
	result = db.Create(allowedAppLanguages)
	if result.Error != nil {
		fmt.Println("Error seeding allowed app languages:", result.Error)
	}
}
