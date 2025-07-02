package models

type MainPricing struct {
	MainPricingId int               `json:"mainPricingId" gorm:"primary_key"`
	StartDate     string            `json:"startDate"`
	EndDate       string            `json:"endDate"`
	Color         string            `json:"color"`
	Name          string            `json:"name"`
	ProductId     int               `json:"-"`
	Pricings      []SpecificPricing `json:"pricings" gorm:"foreignKey:MainPricingId"`
}
