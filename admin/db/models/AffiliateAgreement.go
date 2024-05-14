package models

import (
	"time"

	modelHelpers "github.com/tiqueteo/adminv2-mock-api/db/models/helpers"
	"gorm.io/gorm"
)

type ValueType string

const (
	Value      ValueType = "value"
	Percentage ValueType = "percentage"
)

type AgreementType string

const (
	AgreementTypeTicket   AgreementType = "ticket"
	AgreementTypePurchase AgreementType = "purchase"
)

type AffiliateAgreement struct {
	AffiliateAgreementId int                          `json:"affiliateAgreementId" gorm:"primary_key"`
	Status               bool                         `json:"status"`
	Description          string                       `json:"description"`
	AgreementType        AgreementType                `json:"agreementType"`
	ValueType            ValueType                    `json:"valueType"`
	BaseReturnValue      int                          `json:"baseReturnValue"`
	IsDefault            bool                         `json:"isDefault"`
	WeekDay              modelHelpers.JSONIntSLice    `json:"weekDay" gorm:"type:JSON"`
	DisabledDates        modelHelpers.JSONStringSlice `json:"disabledDates" gorm:"type:JSON"`
	ProductSet           []Product                    `json:"productSet" gorm:"many2many:affiliate_agreement_products;"`
	BuyerTypeSet         []BuyerType                  `json:"buyerTypeSet" gorm:"many2many:affiliate_agreement_buyer_type;"`
	Priority             int                          `json:"priority"`
	CreatedAt            time.Time                    `json:"-"`
	UpdatedAt            time.Time                    `json:"-"`
	DeletedAt            gorm.DeletedAt               `json:"-"`
}
