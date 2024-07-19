package models

import (
	"time"

	"gorm.io/gorm"
)

type VenueCapacity struct {
	VenueCapacityId int            `json:"venueCapacityId" gorm:"primary_key"`
	Name            string         `json:"name"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
