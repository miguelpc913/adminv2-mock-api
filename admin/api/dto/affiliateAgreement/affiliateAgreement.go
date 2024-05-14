package dtoAffiliateAgreement

import "github.com/tiqueteo/adminv2-mock-api/db/models"

type DisplayOrderRequest struct {
	AffiliateAgreementId int `json:"affiliateAgreementId"`
	Priority             int `json:"priority"`
}

type Post struct {
	Status          bool                 `json:"status"`
	Description     string               `json:"description"`
	AgreementType   models.AgreementType `json:"agreementType"`
	ValueType       models.ValueType     `json:"valueType"`
	BaseReturnValue int                  `json:"baseReturnValue"`
	IsDefault       bool                 `json:"isDefault"`
	WeekDay         []int                `json:"weekDay"`
	DisabledDates   []string             `json:"disabledDates"`
	ProductSet      []int                `json:"productSet"`
	BuyerTypeSet    []int                `json:"buyerTypeSet"`
}

type General struct {
	Status          bool                 `json:"status"`
	Description     string               `json:"description"`
	AgreementType   models.AgreementType `json:"agreementType"`
	ValueType       models.ValueType     `json:"valueType"`
	BaseReturnValue int                  `json:"baseReturnValue"`
	IsDefault       bool                 `json:"isDefault"`
}

type Validities struct {
	WeekDay       []int    `json:"weekDay"`
	DisabledDates []string `json:"disabledDates"`
}
