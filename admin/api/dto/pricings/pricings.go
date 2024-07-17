package dto

type ProductVenueBuyerTypeDTO struct {
	ProductVenueBuyerTypeId int     `json:"productVenueBuyerTypeId"`
	ProductId               int     `json:"productId"`
	VenueId                 int     `json:"venueId"`
	BuyerTypeId             int     `json:"buyerTypeId"`
	Price                   float64 `json:"price"`
	HasDiscount             bool    `json:"hasDiscount"`
}

type ProductExtraBuyerTypeDTO struct {
	ProductExtraBuyerTypeId int     `json:"productExtraBuyerTypeId"`
	ProductId               int     `json:"productId"`
	ExtraId                 int     `json:"extraId"`
	BuyerTypeId             int     `json:"buyerTypeId"`
	Price                   float64 `json:"price"`
	HasDiscount             bool    `json:"hasDiscount"`
}

type SpecificPricingDTO struct {
	PricingId              *int                       `json:"pricingId"`
	Name                   string                     `json:"name"`
	Priority               int                        `json:"priority"`
	WeekDays               []int                      `json:"weekDays"`
	EnabledDates           []string                   `json:"enabledDates"`
	ProductVenueBuyerTypes []ProductVenueBuyerTypeDTO `json:"productVenueBuyerTypes"`
	ProductExtraBuyerTypes []ProductExtraBuyerTypeDTO `json:"productExtraBuyerTypes"`
	Default                bool                       `json:"default"`
}

type PricingDTO struct {
	MainPricingId *int                 `json:"mainPricingId"`
	StartDatetime string               `json:"startDatetime"`
	EndDatetime   string               `json:"endDatetime"`
	Pricings      []SpecificPricingDTO `json:"pricings"`
}
