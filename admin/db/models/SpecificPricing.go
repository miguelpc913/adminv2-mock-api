package models

import (
	"time"

	modelHelpers "github.com/tiqueteo/adminv2-mock-api/db/models/helpers"
	"gorm.io/gorm"
)

type SpecificPricing struct {
	PricingId              int                          `json:"pricingId" gorm:"primary_key"`
	MainPricingId          int                          `json:"-"`
	Name                   string                       `json:"name"`
	Priority               int                          `json:"priority"`
	Weekdays               modelHelpers.JSONIntSLice    `json:"weekDays" gorm:"type:JSON"`
	EnabledDates           modelHelpers.JSONStringSlice `json:"enabledDates" gorm:"type:JSON"`
	ProductVenueBuyerTypes []ProductVenueBuyerTypes     `json:"productVenueBuyerTypes" gorm:"foreignKey:PricingId"`
	ProductExtraBuyerTypes []ProductExtraBuyerTypes     `json:"productExtraBuyerTypes" gorm:"foreignKey:PricingId"`
	Default                bool                         `json:"default"`
	CreatedAt              time.Time                    `json:"-"`
	UpdatedAt              time.Time                    `json:"-"`
	DeletedAt              gorm.DeletedAt               `json:"-"`
}
