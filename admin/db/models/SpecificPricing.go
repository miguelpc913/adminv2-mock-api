package models

import (
	"time"

	modelHelpers "github.com/tiqueteo/adminv2-mock-api/db/models/helpers"
	"gorm.io/gorm"
)

type RecurrentTime struct {
	RecurrentTimeId int                       `json:"-" gorm:"primaryKey"`
	PricingId       int                       `json:"-"`
	Minutes         modelHelpers.JSONIntSLice `json:"minutes" gorm:"type:JSON"`
	Hours           modelHelpers.JSONIntSLice `json:"hours" gorm:"type:JSON"`
}

type OccupancyRange struct {
	OccupancyRangeId              int `json:"-" gorm:"primaryKey"`
	DynamicPricingConfigurationId int `json:"-"`
	Start                         int `json:"start"`
	End                           int `json:"end"`
}

type DynamicPricingConfiguration struct {
	DynamicPricingConfigurationId int              `json:"-" gorm:"primaryKey"`
	PricingId                     int              `json:"-"`
	Type                          string           `json:"type"`                                                            // Assuming string. Adjust if it’s an enum or object.
	StartHour                     string           `json:"startHour"`                                                       // Only relevant when type == event_range
	EndHour                       string           `json:"endHour"`                                                         // Only relevant when type == event_range
	OccupancyRanges               []OccupancyRange `json:"occupancyRanges" gorm:"foreignKey:DynamicPricingConfigurationId"` // must be serialized as JSON
}

type SpecificPricing struct {
	PricingId                   int                          `json:"pricingId" gorm:"primaryKey"`
	MainPricingId               int                          `json:"-"`
	Name                        string                       `json:"name"`
	Priority                    int                          `json:"priority"`
	Weekdays                    modelHelpers.JSONIntSLice    `json:"weekDays" gorm:"type:JSON"`
	EnabledDates                modelHelpers.JSONStringSlice `json:"enabledDates" gorm:"type:JSON"`
	StartHour                   *string                      `json:"startHour"`
	EndHour                     *string                      `json:"endHour"`
	RecurrentTime               *RecurrentTime               `json:"recurrentTime" gorm:"foreignKey:PricingId"`
	DynamicPricingConfiguration *DynamicPricingConfiguration `json:"dynamicPricingConfiguration" gorm:"foreignKey:PricingId"`
	ProductVenueBuyerTypes      []ProductVenueBuyerTypes     `json:"productVenueBuyerTypes" gorm:"foreignKey:PricingId"`
	ProductExtraBuyerTypes      []ProductExtraBuyerTypes     `json:"productExtraBuyerTypes" gorm:"foreignKey:PricingId"`
	Default                     bool                         `json:"default"`
	CreatedAt                   time.Time                    `json:"-"`
	UpdatedAt                   time.Time                    `json:"-"`
	DeletedAt                   gorm.DeletedAt               `json:"-"`
}
