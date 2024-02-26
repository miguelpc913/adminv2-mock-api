package models

import (
	"time"

	dbhelpers "github.com/tiqueteo/adminv2-mock-api/db/helpers"
	"gorm.io/gorm"
)

type PromotionType string

const (
	EventType        PromotionType = "event"
	Promotional_code PromotionType = "promotional_code"
	Agreement        PromotionType = "agreement"
	Seller_agreement PromotionType = "seller_agreement"
	Volume           PromotionType = "volume"
)

type RedeemType string

const (
	Ticket      RedeemType = "ticket"
	Reservation RedeemType = "reservation"
)

type CodeType string

const (
	Given     CodeType = "given"
	Generated CodeType = "generated"
)

type Promotion struct {
	PromotionId                 int                       `json:"promotionId" gorm:"primary_key"`
	Status                      bool                      `json:"status"`
	Name                        string                    `json:"name"`
	ShortName                   string                    `json:"shortName"`
	PromotionType               PromotionType             `json:"promotionType"`
	Amount                      int                       `json:"amount"`
	Percentage                  int                       `json:"percentage"`
	LeftPurchased               int                       `json:"leftPurchased"`
	RightPaid                   int                       `json:"rightPaid"`
	IsPromotionAffiliateEnabled bool                      `json:"isPromotionAffiliateEnabled"`
	HideAmountAtTicket          bool                      `json:"hideAmountAtTicket"`
	ShowOriginalAmountAtTicket  bool                      `json:"showOriginalAmountAtTicket"`
	IsGrouped                   bool                      `json:"isGrouped"`
	RedeemType                  RedeemType                `json:"redeemType"`
	CodeType                    CodeType                  `json:"codeType"`
	NumberOfCodes               int                       `json:"numberOfCodes"`
	CodeLength                  int                       `json:"codeLength"`
	Quantity                    int                       `json:"quantity"`
	PromotionalCodeSet          []PromotionalCode         `json:"promotionalCodeSet"`
	StartDatetime               time.Time                 `json:"startDatetime"`
	EndDatetime                 time.Time                 `json:"endDatetime"`
	EventStartDatetime          time.Time                 `json:"eventStartDatetime"`
	EventEndDatetime            time.Time                 `json:"eventEndDatetime"`
	MinSecondsBeforeEvent       int                       `json:"minSecondsBeforeEvent"`
	MaxSecondsBeforeEvent       int                       `json:"maxSecondsBeforeEvent"`
	WeekDay                     dbhelpers.JSONIntSLice    `json:"weekDay" gorm:"type:JSON"`
	DisabledDates               dbhelpers.JSONStringSlice `json:"disabledDates" gorm:"type:JSON"`
	PromotionPriceSet           []PromotionPrice          `json:"promotionPriceSet"`
	CreatedAt                   time.Time                 `json:"-"`
	UpdatedAt                   time.Time                 `json:"-"`
	DeletedAt                   gorm.DeletedAt            `json:"-"`
	SalesGroupSet               []SalesGroup              `json:"salesGroupSet" gorm:"many2many:promotion_salesgroups;"`
	ProductSet                  []Product                 `json:"productSet" gorm:"many2many:promotion_products;"`
	BuyerTypeSet                []BuyerType               `json:"buyerTypeSet" gorm:"many2many:promotion_buyer_type;"`
}
