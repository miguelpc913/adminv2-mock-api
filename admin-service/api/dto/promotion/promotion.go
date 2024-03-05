package dtoPromotion

import "github.com/tiqueteo/adminv2-mock-api/db/models"

type salesGroupPost struct {
	SalesGroupId int    `json:"salesGroupId"`
	Name         string `json:"name"`
}

type productsPost struct {
	ProductId int    `json:"productId"`
	Name      string `json:"name"`
}

type buyerTypePost struct {
	BuyerTypeID int    `json:"buyerTypeId"`
	Name        string `json:"name"`
}

type promotionalCodePost struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type promotionPrice struct {
	SalesGroupId int `json:"salesGroupId"`
	BuyerTypeID  int `json:"buyerTypeId"`
	Amount       int `json:"amount"`
	Percentage   int `json:"percentage"`
}

type PromotionPost struct {
	PromotionId                 int                   `json:"promotionId" gorm:"primary_key"`
	Status                      bool                  `json:"status"`
	Name                        string                `json:"name"`
	ShortName                   string                `json:"shortName"`
	PromotionType               models.PromotionType  `json:"promotionType"`
	Amount                      int                   `json:"amount"`
	Percentage                  int                   `json:"percentage"`
	LeftPurchased               int                   `json:"leftPurchased"`
	RightPaid                   int                   `json:"rightPaid"`
	StartTime                   string                `json:"startTime"`
	EndTime                     string                `json:"endTime"`
	IsPromotionAffiliateEnabled bool                  `json:"isPromotionAffiliateEnabled"`
	HideAmountAtTicket          bool                  `json:"hideAmountAtTicket"`
	ShowOriginalAmountAtTicket  bool                  `json:"showOriginalAmountAtTicket"`
	IsGrouped                   bool                  `json:"isGrouped"`
	RedeemType                  models.RedeemType     `json:"redeemType"`
	CodeType                    models.CodeType       `json:"codeType"`
	NumberOfCodes               int                   `json:"numberOfCodes"`
	CodeLength                  int                   `json:"codeLength"`
	Quantity                    int                   `json:"quantity"`
	PromotionalCodeSet          []promotionalCodePost `json:"promotionalCodeSet"`
	StartDatetime               string                `json:"startDatetime"`
	EndDatetime                 string                `json:"endDatetime"`
	EventStartDatetime          *string               `json:"eventStartDatetime"`
	EventEndDatetime            *string               `json:"eventEndDatetime"`
	MinSecondsBeforeEvent       int                   `json:"minSecondsBeforeEvent"`
	MaxSecondsBeforeEvent       int                   `json:"maxSecondsBeforeEvent"`
	WeekDay                     []int                 `json:"weekDay"`
	DisabledDates               []string              `json:"disabledDates"`
	SalesGroupSet               []salesGroupPost      `json:"salesGroupSet"`
	ProductSet                  []productsPost        `json:"productSet"`
	BuyerTypeSet                []buyerTypePost       `json:"buyerTypeSet"`
}

type PutPromotionIdentity struct {
	Status    bool   `json:"status"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type PutPromotionGeneral struct {
	Amount                      int                   `json:"amount"`
	Percentage                  int                   `json:"percentage"`
	LeftPurchased               int                   `json:"leftPurchased"`
	RightPaid                   int                   `json:"rightPaid"`
	IsPromotionAffiliateEnabled bool                  `json:"isPromotionAffiliateEnabled"`
	HideAmountAtTicket          bool                  `json:"hideAmountAtTicket"`
	ShowOriginalAmountAtTicket  bool                  `json:"showOriginalAmountAtTicket"`
	IsGrouped                   bool                  `json:"isGrouped"`
	RedeemType                  models.RedeemType     `json:"redeemType"`
	CodeType                    models.CodeType       `json:"codeType"`
	NumberOfCodes               int                   `json:"numberOfCodes"`
	CodeLength                  int                   `json:"codeLength"`
	Quantity                    int                   `json:"quantity"`
	PromotionalCodeSet          []promotionalCodePost `json:"promotionalCodeSet"`
}

type PutPromotionValidities struct {
	StartDatetime         string   `json:"startDatetime"`
	EndDatetime           string   `json:"endDatetime"`
	EventStartDatetime    *string  `json:"eventStartDatetime"`
	EventEndDatetime      *string  `json:"eventEndDatetime"`
	StartTime             string   `json:"startTime"`
	EndTime               string   `json:"endTime"`
	MinSecondsBeforeEvent int      `json:"minSecondsBeforeEvent"`
	MaxSecondsBeforeEvent int      `json:"maxSecondsBeforeEvent"`
	WeekDay               []int    `json:"weekDay"`
	DisabledDates         []string `json:"disabledDates"`
}

type PutPromotionAdvancedSettings struct {
	PromotionPriceSet []promotionPrice `json:"promotionPriceSet"`
}

type PutPromotionSalesGroups struct {
	SalesGroupSet []salesGroupPost `json:"salesGroupSet"`
}

type PutPromotionBuyerTypes struct {
	BuyerTypeSet []buyerTypePost `json:"buyerTypeSet"`
}
type PutPromotionProducts struct {
	ProductSet []productsPost `json:"productSet"`
}

type CodeToValidate struct {
	Code string `json:"code"`
}
