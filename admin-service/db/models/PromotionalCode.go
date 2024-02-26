package models

import (
	"time"

	"gorm.io/gorm"
)

type PromotionalCode struct {
	PromotionalCodeId int            `json:"promotionalCodeId" gorm:"primary_key"`
	PromotionId       int            `json:"promotionId"`
	Code              string         `json:"code"`
	Quantity          string         `json:"quantity"`
	CreatedAt         time.Time      `json:"-"`
	UpdatedAt         time.Time      `json:"-"`
	DeletedAt         gorm.DeletedAt `json:"-"`
}
