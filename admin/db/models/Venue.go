package models

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	VenueId   int            `json:"venueId" gorm:"primary_key"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
