package models

import (
	"time"

	"gorm.io/gorm"
)

type PromotionPrice struct {
	PromotionPriceId int            `json:"promotionPriceId" gorm:"primary_key"`
	PromotionId      int            `json:"promotionId"`
	Promotion        Promotion      `json:"-"`
	BuyerTypeId      int            `json:"buyerTypeId"`
	BuyerType        BuyerType      `json:"-"`
	SalesGroupId     int            `json:"salesGroupId"`
	SalesGroup       SalesGroup     `json:"-"`
	Percentage       int            `json:"percentage"`
	Amount           int            `json:"amount"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}
