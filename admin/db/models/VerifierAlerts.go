package models

import (
	"time"

	"gorm.io/gorm"
)

type VerifierAlertPromotion struct {
	VerifierAlertId uint           `gorm:"primaryKey" json:"verifierAlertId"`
	PromotionID     int            `json:"promotionId"`
	Promotion       Promotion      `json:"promotion"`
	AlertColor      string         `json:"alertColor"`
	AlertSound      string         `json:"alertSound"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type VerifierAlertBuyerType struct {
	VerifierAlertId uint           `gorm:"primaryKey" json:"verifierAlertId"`
	BuyerTypeID     int            `json:"buyerTypeId"`
	BuyerType       BuyerType      `json:"buyerType"`
	AlertColor      string         `json:"alertColor"`
	AlertSound      string         `json:"alertSound"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
