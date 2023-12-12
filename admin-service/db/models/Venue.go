package models

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	VenueCapacityId int            `json:"venueCapacityId" gorm:"primary_key"`
	Name            string         `json:"name"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
