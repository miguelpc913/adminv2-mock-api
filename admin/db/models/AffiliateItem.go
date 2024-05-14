package models

import (
	"time"

	"gorm.io/gorm"
)

type ItemType string

const (
	CustomUrl ItemType = "custom_url"
	Banner    ItemType = "banner"
)

type AffiliateItem struct {
	AffiliateItemId int            `json:"affiliateItemId" gorm:"primary_key"`
	Status          bool           `json:"status"`
	ItemName        string         `json:"itemName"`
	ItemType        ItemType       `json:"itemType"`
	ItemLinkedUrl   string         `json:"itemLinkedUrl"`
	ItemResource    *string        `json:"itemResource"`
	IsGeneric       bool           `json:"isGeneric"`
	ProductSet      []Product      `json:"productSet" gorm:"many2many:affiliate_item_products;"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
