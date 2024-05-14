package models

import (
	"time"

	modelHelpers "github.com/tiqueteo/adminv2-mock-api/db/models/helpers"
	"gorm.io/gorm"
)

type ProductInfo struct {
	ProductInfoId     int                          `json:"productInfoId" gorm:"primary_key"`
	Status            bool                         `json:"status"`
	Name              string                       `json:"name"`
	InternalName      string                       `json:"internalName"`
	Description       string                       `json:"description"`
	Icon              string                       `json:"icon"`
	ProductInfoTypeId int                          `json:"-"`
	ProductInfoType   ProductInfoType              `json:"-"`
	InfoType          string                       `json:"infoType" gorm:"-"`
	CalendarColor     string                       `json:"calendarColor"`
	Weekdays          modelHelpers.JSONIntSLice    `json:"weekDays" gorm:"type:JSON"`
	SelectedDates     modelHelpers.JSONStringSlice `json:"selectedDates" gorm:"type:JSON"`
	DisplayOrder      int                          `json:"displayOrder"`
	CreatedAt         time.Time                    `json:"-"`
	UpdatedAt         time.Time                    `json:"-"`
	DeletedAt         gorm.DeletedAt               `json:"-"`
	SalesGroupSet     []SalesGroup                 `json:"salesGroupSet" gorm:"many2many:productinfo_salesgroups;"`
	ProductSet        []Product                    `json:"productSet" gorm:"many2many:productinfo_products;"`
	VenueSet          []Venue                      `json:"venueSet" gorm:"many2many:productinfo_venue;"`
}
