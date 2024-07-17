package models

type ProductExtraBuyerTypes struct {
	ProductExtraBuyerTypeId int     `gorm:"primary_key" json:"productExtraBuyerTypeId"`
	ProductId               int     `json:"productId"`
	ExtraId                 int     `json:"extraId"`
	PricingId               int     `json:"pricingId"`
	BuyerTypeId             int     `json:"buyerTypeId"`
	Price                   float64 `json:"price"`
	HasDiscount             bool    `json:"hasDiscount"`
}
