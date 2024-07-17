package models

type ProductVenueBuyerTypes struct {
	ProductVenueBuyerTypeId int     `gorm:"primary_key" json:"productVenueBuyerTypeId"`
	ProductId               int     `json:"productId"`
	PricingId               int     `json:"pricingId"`
	VenueId                 int     `json:"venueId"`
	BuyerTypeId             int     `json:"buyerTypeId"`
	Price                   float64 `json:"price"`
	HasDiscount             bool    `json:"hasDiscount"`
}
