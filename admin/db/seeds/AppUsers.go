package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	// Create PointsOfSales
	posSets := []models.PointsOfSale{
		{Name: "POS Set A"},
		{Name: "POS Set B"},
	}
	if err := db.Create(&posSets).Error; err != nil {
		fmt.Println("❌ Failed to seed PointsOfSales:", err)
		return
	}

	// Create ReportSets
	reportSets := []models.Report{
		{Name: "Report Set X"},
		{Name: "Report Set Y"},
	}
	if err := db.Create(&reportSets).Error; err != nil {
		fmt.Println("❌ Failed to seed ReportSets:", err)
		return
	}

	// Create Users
	users := []models.User{
		{
			Status:         true,
			Type:           models.UserTypeCallCenter,
			Profile:        models.UserProfileOperator,
			Name:           "Alice",
			LastName:       "Smith",
			UserName:       "alice.smith",
			Email:          "alice@example.com",
			APIUser:        false,
			Password:       "hashed_password_1",
			PointOfSaleSet: []models.PointsOfSale{posSets[0]},
			ReportSet:      []models.Report{reportSets[0]},
		},
		{
			Status:         true,
			Type:           models.UserTypeBoxOffice,
			Profile:        models.UserProfileAdmin,
			Name:           "Bob",
			LastName:       "Jones",
			UserName:       "bob.jones",
			Email:          "bob@example.com",
			APIUser:        true,
			Password:       "hashed_password_2",
			PointOfSaleSet: []models.PointsOfSale{posSets[1]},
			ReportSet:      []models.Report{reportSets[1]},
		},
		{
			Status:         false,
			Type:           models.UserTypeCallCenter,
			Profile:        models.UserProfileAdmin,
			Name:           "Carol",
			LastName:       "White",
			UserName:       "carol.white",
			Email:          "carol@example.com",
			APIUser:        false,
			Password:       "hashed_password_3",
			PointOfSaleSet: posSets,    // all POS sets
			ReportSet:      reportSets, // all Report sets
		},
	}

	if err := db.Create(&users).Error; err != nil {
		fmt.Println("❌ Failed to seed Users:", err)
	}
}
