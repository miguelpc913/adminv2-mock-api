package models

import (
	"time"

	"gorm.io/gorm"
)

type Tax struct {
	TaxId     int            `json:"taxId"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Extra struct {
	ExtraId          int            `json:"extraId" gorm:"primary_key"`
	Name             string         `json:"name"`
	Status           bool           `json:"status"`
	TaxId            int            `json:"-"`
	Tax              Tax            `json:"tax" gorm:"foreignKey:TaxId"`
	HideFromSummary  bool           `json:"hideFromSummary"`
	ManagesStock     bool           `json:"managesStock"`
	ManagesStockType string         `json:"managesStockType"`
	Stock            int            `json:"stock"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}
