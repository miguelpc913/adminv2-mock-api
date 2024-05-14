package seeds

import (
	"fmt"
	"time"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedAffiliateAgreements(db *gorm.DB) {
	affiliateAgreements := []models.AffiliateAgreement{
		{
			Status:          true,
			Description:     "Convenio de Comisión Básica",
			AgreementType:   "ticket",
			ValueType:       "value",
			BaseReturnValue: 100,
			IsDefault:       true,
			WeekDay:         []int{1, 2, 3, 4, 5},
			DisabledDates:   []string{"2024-02-15", "2024-02-20"},
			Priority:        1,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DeletedAt:       gorm.DeletedAt{},
		},
		{
			Status:          true,
			Description:     "Convenio de Comisión Alta",
			AgreementType:   "purchase",
			ValueType:       "percentage",
			BaseReturnValue: 10,
			IsDefault:       false,
			WeekDay:         []int{1, 3, 5},
			DisabledDates:   []string{"2024-02-18", "2024-02-25"},
			Priority:        2,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DeletedAt:       gorm.DeletedAt{},
		},
		{
			Status:          true,
			Description:     "Convenio de Comisión Premium",
			AgreementType:   "ticket",
			ValueType:       "value",
			BaseReturnValue: 200,
			IsDefault:       false,
			WeekDay:         []int{2, 4, 6},
			DisabledDates:   []string{"2024-02-21", "2024-02-28"},
			Priority:        3,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DeletedAt:       gorm.DeletedAt{},
		},
		{
			Status:          true,
			Description:     "Convenio de Comisión Preferencial",
			AgreementType:   "purchase",
			ValueType:       "value",
			BaseReturnValue: 50,
			IsDefault:       true,
			WeekDay:         []int{3, 5, 7},
			DisabledDates:   []string{"2024-02-22", "2024-02-29"},
			Priority:        4,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DeletedAt:       gorm.DeletedAt{},
		},
		{
			Status:          true,
			Description:     "Convenio de Comisión Plus",
			AgreementType:   "ticket",
			ValueType:       "percentage",
			BaseReturnValue: 15,
			IsDefault:       true,
			WeekDay:         []int{4, 6, 1},
			DisabledDates:   []string{"2024-02-23", "2024-03-01"},
			Priority:        5,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			DeletedAt:       gorm.DeletedAt{},
		},
	}

	result := db.Create(&affiliateAgreements)
	if result.Error != nil {
		fmt.Println("Error seeding affiliate_agreements:", result.Error)
	}
}
