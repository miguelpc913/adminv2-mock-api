package models

import (
	"time"

	"gorm.io/gorm"
)

type AffiliateStatus string

const (
	PendingAffiliate  AffiliateStatus = "pending"
	EnabledAffiliate  AffiliateStatus = "enabled"
	DisabledAffiliate AffiliateStatus = "disabled"
	RejectedAffiliate AffiliateStatus = "rejected"
)

type Affiliate struct {
	AffiliateId           int                  `json:"affiliateId" gorm:"primary_key"`
	Status                AffiliateStatus      `json:"status"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	Website               string               `json:"website"`
	ContactEmail          string               `json:"contactEmail"`
	VatNumberType         string               `json:"vatNumberType"`
	VatNumber             string               `json:"vatNumber"`
	ExternalAffiliateId   string               `json:"externalAffiliateId"`
	MinimumPayout         int                  `json:"minimumPayout"`
	AffiliateAgreementSet []AffiliateAgreement `json:"affiliateAgreementSet" gorm:"many2many:affiliate_affiliate_agreement;"`
	CreatedAt             time.Time            `json:"-"`
	UpdatedAt             time.Time            `json:"-"`
	DeletedAt             gorm.DeletedAt       `json:"-"`
}
