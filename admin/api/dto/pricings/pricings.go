package dto

// type ProductVenueBuyerTypeDTO struct {
// 	ProductVenueBuyerTypeId int     `json:"productVenueBuyerTypeId"`
// 	ProductId               int     `json:"productId"`
// 	VenueId                 int     `json:"venueId"`
// 	BuyerTypeId             int     `json:"buyerTypeId"`
// 	Price                   float64 `json:"price"`
// 	HasDiscount             bool    `json:"hasDiscount"`
// }

// type ProductExtraBuyerTypeDTO struct {
// 	ProductExtraBuyerTypeId int     `json:"productExtraBuyerTypeId"`
// 	ProductId               int     `json:"productId"`
// 	ExtraId                 int     `json:"extraId"`
// 	BuyerTypeId             int     `json:"buyerTypeId"`
// 	Price                   float64 `json:"price"`
// 	HasDiscount             bool    `json:"hasDiscount"`
// }

// type SpecificPricing struct {
// 	PricingId              *int                       `json:"pricingId"`
// 	Name                   string                     `json:"name"`
// 	Priority               int                        `json:"priority"`
// 	WeekDays               []int                      `json:"weekDays"`
// 	EnabledDates           []string                   `json:"enabledDates"`
// 	ProductVenueBuyerTypes []ProductVenueBuyerTypeDTO `json:"productVenueBuyerTypes"`
// 	ProductExtraBuyerTypes []ProductExtraBuyerTypeDTO `json:"productExtraBuyerTypes"`
// 	Default                bool                       `json:"default"`
// }

// type PricingDTO struct {
// 	MainPricingId *int              `json:"mainPricingId"`
// 	StartDatetime string            `json:"startDatetime"`
// 	EndDatetime   string            `json:"endDatetime"`
// 	Pricings      []SpecificPricing `json:"pricings"`
// }

type UpdatePricingValuesDTO []PricingUpdateDTO

type PricingUpdateDTO struct {
	PricingId              int                          `json:"pricingId" binding:"required"`
	ProductVenueBuyerTypes []ProductVenueBuyerTypePrice `json:"productVenueBuyerTypes,omitempty"`
	ProductExtraBuyerTypes []ProductExtraBuyerTypePrice `json:"productExtraBuyerTypes,omitempty"`
}

type ProductVenueBuyerTypePrice struct {
	ProductVenueBuyerTypeId int     `json:"productVenueBuyerTypeId" binding:"required"`
	Price                   float64 `json:"price" binding:"required"`
}

type ProductExtraBuyerTypePrice struct {
	ProductExtraBuyerTypeId int     `json:"productExtraBuyerTypeId" binding:"required"`
	Price                   float64 `json:"price" binding:"required"`
}

type UpdatePricingPriorityDTO []PriorityUpdateDTO

type PriorityUpdateDTO struct {
	PricingId int `json:"pricingId" binding:"required"`
	Priority  int `json:"priority"`
}

type BasePricingPost struct {
	Color     string `json:"color"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type SpecificPricingDTO struct {
	Name                        string                          `json:"name"`
	WeekDays                    []int                           `json:"weekDays,omitempty"`                    // optional
	EnabledDates                []string                        `json:"enabledDates,omitempty"`                // optional, format: YYYY-MM-DD
	StartHour                   *string                         `json:"startHour,omitempty"`                   // optional, format: time string
	EndHour                     *string                         `json:"endHour,omitempty"`                     // optional
	RecurrentTime               *RecurrentTimeDTO               `json:"recurrentTime,omitempty"`               // optional
	DynamicPricingConfiguration *DynamicPricingConfigurationDTO `json:"dynamicPricingConfiguration,omitempty"` // optional
	Default                     bool                            `json:"default,omitempty"`                     // optional
}

type RecurrentTimeDTO struct {
	Minutes []int `json:"minutes"` // optional
	Hours   []int `json:"hours"`   // optional
}

type DynamicPricingConfigurationDTO struct {
	Type            string              `json:"type"` // optional but assumed string default
	StartHour       *string             `json:"startHour,omitempty"`
	EndHour         *string             `json:"endHour,omitempty"`
	OccupancyRanges []OccupancyRangeDTO `json:"occupancyRanges"`
}

type OccupancyRangeDTO struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
