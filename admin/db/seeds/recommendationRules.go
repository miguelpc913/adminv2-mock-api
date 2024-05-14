package seeds

import (
	"fmt"
	"time"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedRecommendationRules(db *gorm.DB) {
	recommendationRules := []models.RecommendationRule{
		{
			Status:                     true,
			Name:                       `Summer Special`,
			OfferingType:               `Online`,
			ProductId:                  1,
			OfferedProductId:           2,
			DirectAddToCart:            true,
			StartDatetime:              time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			EndDatetime:                time.Date(2023, 8, 31, 23, 59, 59, 0, time.UTC),
			EventStartDatetime:         time.Date(2023, 6, 15, 8, 0, 0, 0, time.UTC),
			EventEndDatetime:           time.Date(2023, 8, 15, 17, 0, 0, 0, time.UTC),
			WeekDay:                    []int{1, 3, 5},
			StartTime:                  `09:00`,
			EndTime:                    `18:00`,
			SessionOffsetMinutesBefore: 30,
			SessionOffsetMinutesAfter:  15,
			Title:                      `{"en": "Summer Sale", "es": "Rebajas de Verano"}`,
			Body:                       `{"en": "Get ready for summer with special discounts!", "es": "¡Prepárate para el verano con descuentos especiales!"}`,
			Footer:                     `{"en": "Limited time offer", "es": "Oferta por tiempo limitado"}`,
			Priority:                   1,
		},
		{
			Status:                     true,
			Name:                       `Winter Collection`,
			OfferingType:               `Online`,
			ProductId:                  1,
			OfferedProductId:           2,
			DirectAddToCart:            false,
			StartDatetime:              time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
			EndDatetime:                time.Date(2024, 2, 28, 23, 59, 59, 0, time.UTC),
			EventStartDatetime:         time.Date(2023, 12, 1, 9, 0, 0, 0, time.UTC),
			EventEndDatetime:           time.Date(2024, 2, 1, 18, 0, 0, 0, time.UTC),
			WeekDay:                    []int{2, 4, 6},
			StartTime:                  `10:00`,
			EndTime:                    `19:00`,
			SessionOffsetMinutesBefore: 45,
			SessionOffsetMinutesAfter:  20,
			Title:                      `{"en": "Winter Collection Launch", "es": "Lanzamiento de la Colección de Invierno"}`,
			Body:                       `{"en": "Explore our new winter collection!", "es": "¡Explora nuestra nueva colección de invierno!"}`,
			Footer:                     `{"en": "Stay warm and stylish", "es": "Mantente cálido y con estilo"}`,
			Priority:                   2,
		},
		{
			Status:                     true,
			Name:                       `Back to School`,
			OfferingType:               `Online`,
			ProductId:                  1,
			OfferedProductId:           2,
			DirectAddToCart:            true,
			StartDatetime:              time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			EndDatetime:                time.Date(2023, 9, 30, 23, 59, 59, 0, time.UTC),
			EventStartDatetime:         time.Date(2023, 8, 10, 10, 0, 0, 0, time.UTC),
			EventEndDatetime:           time.Date(2023, 9, 10, 19, 0, 0, 0, time.UTC),
			WeekDay:                    []int{1, 2, 3},
			StartTime:                  `11:00`,
			EndTime:                    `20:00`,
			SessionOffsetMinutesBefore: 60,
			SessionOffsetMinutesAfter:  25,
			Title:                      `{"en": "Back to School Deals", "es": "Ofertas de Vuelta al Colegio"}`,
			Body:                       `{"en": "Special offers for the new academic year", "es": "Ofertas especiales para el nuevo año académico"}`,
			Footer:                     `{"en": "Equip for success", "es": "Equípate para el éxito"}`,
			Priority:                   3,
		},
		{
			Status:                     true,
			Name:                       `Fitness Frenzy`,
			OfferingType:               `Online`,
			ProductId:                  1,
			OfferedProductId:           2,
			DirectAddToCart:            false,
			StartDatetime:              time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDatetime:                time.Date(2023, 3, 31, 23, 59, 59, 0, time.UTC),
			EventStartDatetime:         time.Date(2023, 1, 10, 11, 0, 0, 0, time.UTC),
			EventEndDatetime:           time.Date(2023, 3, 10, 20, 0, 0, 0, time.UTC),
			WeekDay:                    []int{4, 5, 6},
			StartTime:                  `12:00`,
			EndTime:                    `21:00`,
			SessionOffsetMinutesBefore: 75,
			SessionOffsetMinutesAfter:  30,
			Title:                      `{"en": "New Year Fitness Goals", "es": "Metas de Fitness para el Año Nuevo"}`,
			Body:                       `{"en": "Kickstart your fitness journey", "es": "Inicia tu viaje de fitness"}`,
			Footer:                     `{"en": "Healthy living starts now", "es": "La vida saludable comienza ahora"}`,
			Priority:                   4,
		},
		{
			Status:                     true,
			Name:                       `Tech Gadgets Galore`,
			OfferingType:               `Online`,
			ProductId:                  1,
			OfferedProductId:           2,
			DirectAddToCart:            true,
			StartDatetime:              time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDatetime:                time.Date(2023, 6, 30, 23, 59, 59, 0, time.UTC),
			EventStartDatetime:         time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC),
			EventEndDatetime:           time.Date(2023, 6, 15, 21, 0, 0, 0, time.UTC),
			WeekDay:                    []int{1, 3, 5},
			StartTime:                  `13:00`,
			EndTime:                    `22:00`,
			SessionOffsetMinutesBefore: 90,
			SessionOffsetMinutesAfter:  35,
			Title:                      `{"en": "Latest in Tech", "es": "Lo Último en Tecnología"}`,
			Body:                       `{"en": "Discover the newest tech gadgets", "es": "Descubre los más recientes gadgets tecnológicos"}`,
			Footer:                     `{"en": "Innovate your life", "es": "Innova tu vida"}`,
			Priority:                   5,
		},
	}

	result := db.Create(&recommendationRules)
	if result.Error != nil {
		fmt.Println("Error occurred while seeding recommendation rules:", result.Error)
	}
}
