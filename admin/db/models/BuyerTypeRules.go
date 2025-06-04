package models

import (
	"time"

	modelHelpers "github.com/tiqueteo/adminv2-mock-api/db/models/helpers"
	"gorm.io/gorm"
)

type Vars struct {
	VarId           uint                      `gorm:"primaryKey"`
	X               modelHelpers.JSONIntSLice `json:"x" gorm:"type:JSON"`
	Y               modelHelpers.JSONIntSLice `json:"y,omitempty" gorm:"type:JSON"`
	M               *int                      `json:"m,omitempty"`
	N               *int                      `json:"n,omitempty"`
	BuyerTypeRuleID uint
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}

type BuyerTypeRule struct {
	BuyerTypeRuleID         uint           `gorm:"primaryKey" json:"buyerTypeRuleId"`
	Status                  bool           `json:"status"`
	Name                    string         `json:"name"`
	BuyerTypeRuleTemplateID uint           `json:"buyerTypeRuleTemplateId"`
	Vars                    Vars           `json:"vars"`
	ErrorMessage            string         `json:"errorMessage"`
	Priority                int            `json:"priority"`
	ProductSet              []Product      `json:"productSet" gorm:"many2many:buyer_type_rules_products;"`
	CreatedAt               time.Time      `json:"-"`
	UpdatedAt               time.Time      `json:"-"`
	DeletedAt               gorm.DeletedAt `json:"-"`
}
