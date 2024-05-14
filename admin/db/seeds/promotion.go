package seeds

import (
	"fmt"
	"time"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedPromotions(db *gorm.DB) {
	promotions := []models.Promotion{
		{
			Status:                      true,
			Name:                        `{"en": "Promotion Name in English", "es": "Nombre de la promoción en español", "ca": "Nom de la promoció en català"}`,
			ShortName:                   "Short Name 1",
			PromotionType:               models.EventType,
			Amount:                      intPtr(100),
			Percentage:                  intPtr(10),
			LeftPurchased:               intPtr(50),
			RightPaid:                   intPtr(20),
			IsPromotionAffiliateEnabled: boolPtr(true),
			RedeemType:                  redeemTypePtr(models.Ticket),
			HideAmountAtTicket:          false,
			ShowOriginalAmountAtTicket:  true,
			IsGrouped:                   true,
			NumberOfCodes:               intPtr(50),
			CodeLength:                  intPtr(8),
			Quantity:                    intPtr(200),
			StartDatetime:               time.Date(2024, 2, 14, 12, 0, 0, 0, time.UTC),
			EndDatetime:                 time.Date(2024, 2, 28, 12, 0, 0, 0, time.UTC),
			EventStartDatetime:          timePtr(time.Date(2024, 2, 14, 8, 0, 0, 0, time.UTC)),
			EventEndDatetime:            timePtr(time.Date(2024, 2, 28, 20, 0, 0, 0, time.UTC)),
			MinSecondsBeforeEvent:       intPtr(3600),
			MaxSecondsBeforeEvent:       intPtr(7200),
			WeekDay:                     []int{1, 3, 5},
			DisabledDates:               []string{`"2024-02-15"`, `"2024-02-20"`},
			StartTime:                   stringPtr("09:00"),
			EndTime:                     stringPtr("10:00"),
			CreatedAt:                   time.Now(),
			UpdatedAt:                   time.Now(),
		},
		{
			Status:                      false,
			Name:                        `{"en": "Second Promotion Name", "es": "Segundo nombre de la promoción", "ca": "Segon nom de la promoció"}`,
			ShortName:                   "Short Name 2",
			PromotionType:               models.PromotionType("promotional_code"),
			Amount:                      intPtr(150),
			Percentage:                  intPtr(15),
			LeftPurchased:               intPtr(75),
			RightPaid:                   intPtr(30),
			IsPromotionAffiliateEnabled: boolPtr(false),
			RedeemType:                  redeemTypePtr(models.Ticket),
			CodeType:                    codeTypePtr(models.CodeType("generated")),
			HideAmountAtTicket:          true,
			ShowOriginalAmountAtTicket:  false,
			IsGrouped:                   false,
			NumberOfCodes:               intPtr(75),
			CodeLength:                  intPtr(6),
			Quantity:                    intPtr(300),
			StartDatetime:               time.Date(2024, 3, 1, 12, 0, 0, 0, time.UTC),
			EndDatetime:                 time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC),
			EventStartDatetime:          timePtr(time.Date(2024, 3, 1, 10, 0, 0, 0, time.UTC)),
			EventEndDatetime:            timePtr(time.Date(2024, 3, 15, 18, 0, 0, 0, time.UTC)),
			MinSecondsBeforeEvent:       intPtr(5400),
			MaxSecondsBeforeEvent:       intPtr(9000),
			WeekDay:                     []int{2, 4, 6},
			DisabledDates:               []string{`"2024-03-02"`, `"2024-03-10"`},
			StartTime:                   stringPtr("09:00"),
			EndTime:                     stringPtr("10:00"),
			CreatedAt:                   time.Now(),
			UpdatedAt:                   time.Now(),
		},
		{
			Status:                      true,
			Name:                        `{"en": "Third Promotion Name", "es": "Tercer nombre de la promoción", "ca": "Tercer nom de la promoció"}`,
			ShortName:                   "Short Name 3",
			PromotionType:               models.Agreement,
			Amount:                      intPtr(200),
			Percentage:                  intPtr(20),
			LeftPurchased:               intPtr(100),
			RightPaid:                   intPtr(40),
			IsPromotionAffiliateEnabled: boolPtr(true),
			RedeemType:                  redeemTypePtr(models.Ticket),
			HideAmountAtTicket:          false,
			ShowOriginalAmountAtTicket:  true,
			IsGrouped:                   true,
			NumberOfCodes:               intPtr(100),
			CodeLength:                  intPtr(10),
			Quantity:                    intPtr(400),
			StartDatetime:               time.Date(2024, 3, 16, 12, 0, 0, 0, time.UTC),
			EndDatetime:                 time.Date(2024, 3, 31, 12, 0, 0, 0, time.UTC),
			EventStartDatetime:          timePtr(time.Date(2024, 3, 16, 9, 0, 0, 0, time.UTC)),
			EventEndDatetime:            timePtr(time.Date(2024, 3, 31, 21, 0, 0, 0, time.UTC)),
			MinSecondsBeforeEvent:       intPtr(7200),
			MaxSecondsBeforeEvent:       intPtr(10800),
			WeekDay:                     []int{1, 5, 7},
			DisabledDates:               []string{`"2024-03-17"`, `"2024-03-25"`},
			StartTime:                   stringPtr("09:00"),
			EndTime:                     stringPtr("10:00"),
			CreatedAt:                   time.Now(),
			UpdatedAt:                   time.Now(),
		},
		{
			Status:                      true,
			Name:                        `{"en": "Fourth Promotion Name", "es": "Cuarto nombre de la promoción", "ca": "Quart nom de la promoció"}`,
			ShortName:                   "Short Name 4",
			PromotionType:               models.Seller_agreement,
			Amount:                      intPtr(250),
			Percentage:                  intPtr(25),
			LeftPurchased:               intPtr(125),
			RightPaid:                   intPtr(50),
			IsPromotionAffiliateEnabled: boolPtr(false),
			RedeemType:                  redeemTypePtr(models.Ticket),
			HideAmountAtTicket:          true,
			ShowOriginalAmountAtTicket:  false,
			IsGrouped:                   false,
			NumberOfCodes:               intPtr(125),
			CodeLength:                  intPtr(12),
			Quantity:                    intPtr(500),
			StartDatetime:               time.Date(2024, 4, 1, 12, 0, 0, 0, time.UTC),
			EndDatetime:                 time.Date(2024, 4, 15, 12, 0, 0, 0, time.UTC),
			EventStartDatetime:          timePtr(time.Date(2024, 4, 1, 11, 0, 0, 0, time.UTC)),
			EventEndDatetime:            timePtr(time.Date(2024, 4, 15, 19, 0, 0, 0, time.UTC)),
			MinSecondsBeforeEvent:       intPtr(9000),
			MaxSecondsBeforeEvent:       intPtr(14400),
			WeekDay:                     []int{2, 6, 1},
			DisabledDates:               []string{`"2024-04-02"`, `"2024-04-10"`},
			StartTime:                   stringPtr("09:00"),
			EndTime:                     stringPtr("10:00"),
			CreatedAt:                   time.Now(),
			UpdatedAt:                   time.Now(),
		},
		{
			Status:                      false,
			Name:                        `{"en": "Fifth Promotion Name", "es": "Quinto nombre de la promoción", "ca": "Cinquè nom de la promoció"}`,
			ShortName:                   "Short Name 5",
			PromotionType:               models.Volume,
			Amount:                      intPtr(300),
			Percentage:                  intPtr(30),
			LeftPurchased:               intPtr(150),
			RightPaid:                   intPtr(60),
			IsPromotionAffiliateEnabled: boolPtr(true),
			RedeemType:                  redeemTypePtr(models.Ticket),
			HideAmountAtTicket:          false,
			ShowOriginalAmountAtTicket:  true,
			IsGrouped:                   true,
			NumberOfCodes:               intPtr(150),
			CodeLength:                  intPtr(14),
			Quantity:                    intPtr(600),
			StartDatetime:               time.Date(2024, 4, 16, 12, 0, 0, 0, time.UTC),
			EndDatetime:                 time.Date(2024, 4, 30, 12, 0, 0, 0, time.UTC),
			EventStartDatetime:          timePtr(time.Date(2024, 4, 16, 8, 0, 0, 0, time.UTC)),
			EventEndDatetime:            timePtr(time.Date(2024, 4, 30, 22, 0, 0, 0, time.UTC)),
			MinSecondsBeforeEvent:       intPtr(12600),
			MaxSecondsBeforeEvent:       intPtr(18000),
			WeekDay:                     []int{3, 7, 2},
			DisabledDates:               []string{`"2024-04-17"`, `"2024-04-25"`},
			StartTime:                   stringPtr("09:00"),
			EndTime:                     stringPtr("10:00"),
			CreatedAt:                   time.Now(),
			UpdatedAt:                   time.Now(),
		},
	}

	result := db.Create(&promotions)
	if result.Error != nil {
		fmt.Println("Error occurred while seeding promotions:", result.Error)
	}
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func redeemTypePtr(r models.RedeemType) *models.RedeemType {
	return &r
}

func codeTypePtr(c models.CodeType) *models.CodeType {
	return &c
}
