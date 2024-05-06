package models

import (
	"time"

	"gorm.io/gorm"
)

type TypeOfVerifier string

const (
	Exit           TypeOfVerifier = "exit"
	ExitValidation TypeOfVerifier = "exit_validation"
	Entry          TypeOfVerifier = "entry"
)

type Verifier struct {
	VerifierId       int            `json:"verifierId" gorm:"primary_key"`
	Name             string         `json:"name"`
	DeviceType       string         `json:"deviceType"`
	TypeOfVerifier   TypeOfVerifier `json:"typeOfVerifier"`
	SerialNumber     string         `json:"serialNumber"`
	AppVersion       string         `json:"appVersion,omitempty"`
	LowerMarginTime  int            `json:"lowerMarginTime"`
	UpperMarginTime  int            `json:"upperMarginTime"`
	DateMode         bool           `json:"dateMode"`
	FixedDate        *time.Time     `json:"fixedDate,omitempty"`
	CheckMaxCapacity bool           `json:"checkMaxCapacity"`
	SalesGroupSet    []SalesGroup   `json:"salesGroupSet" gorm:"many2many:salesgroups_verifiers;"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}
