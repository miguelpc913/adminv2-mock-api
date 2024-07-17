package models

type MainPricing struct {
	MainPricingId int               `json:"mainPricingId" gorm:"primary_key"`
	StartDatetime string            `json:"startDatetime"`
	EndDateTime   string            `json:"endDateTime"`
	ProductId     int               `json:"-"`
	Pricings      []SpecificPricing `json:"pricings" gorm:"foreignKey:MainPricingId"`
}
