package seeds

import (
	"fmt"

	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func SeedVerifier(db *gorm.DB) {
	verifiers := []models.Verifier{
		{
			Name:             "Verifier A",
			DeviceType:       "Type A",
			TypeOfVerifier:   "exit",
			SerialNumber:     "SN-A1",
			AppVersion:       "1.0.1",
			LowerMarginTime:  10,
			UpperMarginTime:  20,
			DateMode:         false,
			CheckMaxCapacity: true,
		},
		{
			Name:             "Verifier B",
			DeviceType:       "Type B",
			TypeOfVerifier:   "exit",
			SerialNumber:     "SN-B2",
			AppVersion:       "1.0.2",
			LowerMarginTime:  15,
			UpperMarginTime:  25,
			DateMode:         true,
			CheckMaxCapacity: false,
		},
		{
			Name:             "Verifier C",
			DeviceType:       "Type C",
			TypeOfVerifier:   "exit",
			SerialNumber:     "SN-C3",
			AppVersion:       "1.0.3",
			LowerMarginTime:  20,
			UpperMarginTime:  30,
			DateMode:         false,
			CheckMaxCapacity: true,
		},
		{
			Name:             "Verifier D",
			DeviceType:       "Type D",
			TypeOfVerifier:   "exit",
			SerialNumber:     "SN-D4",
			AppVersion:       "1.0.4",
			LowerMarginTime:  25,
			UpperMarginTime:  35,
			DateMode:         true,
			CheckMaxCapacity: false,
		},
		{
			Name:             "Verifier E",
			DeviceType:       "Type E",
			TypeOfVerifier:   "exit",
			SerialNumber:     "SN-E5",
			AppVersion:       "1.0.5",
			LowerMarginTime:  30,
			UpperMarginTime:  40,
			DateMode:         false,
			CheckMaxCapacity: true,
		},
	}

	result := db.Create(verifiers)
	if result.Error != nil {
		fmt.Print("There has been an error in the seed of verifier")
	}
}
