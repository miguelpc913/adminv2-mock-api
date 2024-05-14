package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedProductInfos(db *gorm.DB) {
	productInfos := []models.ProductInfo{
		{
			Status:            true,
			Name:              `{"en": "Summer Festival Tickets", "es": "Entradas para el Festival de Verano"}`,
			InternalName:      "summer_fest_tickets",
			Description:       `{"en": "Tickets for the annual Summer Music Festival", "es": "Entradas para el Festival Anual de Música de Verano"}`,
			Icon:              "festival_icon.png",
			CalendarColor:     "#FFD700",
			Weekdays:          []int{1, 2, 3},
			SelectedDates:     []string{"2023-06-21", "2023-06-22"},
			DisplayOrder:      1,
			ProductInfoTypeId: 1,
		},
		{
			Status:            false,
			Name:              `{"en": "Art Exhibition Entry", "es": "Entrada para Exposición de Arte"}`,
			InternalName:      "art_exhibit_entry",
			Description:       `{"en": "Passes for the Modern Art Exhibition", "es": "Pases para la Exposición de Arte Moderno"}`,
			Icon:              "art_icon.png",
			CalendarColor:     "#E6E6FA",
			Weekdays:          []int{4, 5},
			SelectedDates:     []string{"2023-07-15", "2023-07-16"},
			DisplayOrder:      2,
			ProductInfoTypeId: 1,
		},
		{
			Status:            true,
			Name:              `{"en": "Theater Play - The Great Gatsby", "es": "Obra de Teatro - El Gran Gatsby"}`,
			InternalName:      "gatsby_play",
			Description:       `{"en": "Tickets for the play adaptation of The Great Gatsby", "es": "Entradas para la adaptación teatral de El Gran Gatsby"}`,
			Icon:              "theater_icon.png",
			CalendarColor:     "#B0C4DE",
			Weekdays:          []int{6, 0},
			SelectedDates:     []string{"2023-08-05", "2023-08-06"},
			DisplayOrder:      3,
			ProductInfoTypeId: 1,
		},
		{
			Status:            true,
			Name:              `{"en": "Wine Tasting Event", "es": "Evento de Cata de Vinos"}`,
			InternalName:      "wine_tasting",
			Description:       `{"en": "Exclusive wine tasting event tickets", "es": "Entradas para evento exclusivo de cata de vinos"}`,
			Icon:              "wine_icon.png",
			CalendarColor:     "#F5F5DC",
			Weekdays:          []int{1, 3, 5},
			SelectedDates:     []string{"2023-09-10", "2023-09-11"},
			DisplayOrder:      4,
			ProductInfoTypeId: 1,
		},
		{
			Status:            false,
			Name:              `{"en": "Cooking Workshop", "es": "Taller de Cocina"}`,
			InternalName:      "cooking_workshop",
			Description:       `{"en": "Join our Italian cooking workshop", "es": "Únete a nuestro taller de cocina italiana"}`,
			Icon:              "cooking_icon.png",
			CalendarColor:     "#FFFACD",
			Weekdays:          []int{2, 4},
			SelectedDates:     []string{"2023-10-20", "2023-10-21"},
			DisplayOrder:      5,
			ProductInfoTypeId: 1,
		},
	}

	result := db.Create(&productInfos)
	if result.Error != nil {
		fmt.Println("Error occurred while seeding product infos:", result.Error)
	}
}
