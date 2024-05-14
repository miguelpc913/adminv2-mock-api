package dtoAffiliateItem

import "github.com/tiqueteo/adminv2-mock-api/db/models"

type AffiliateItemReq struct {
	Status        bool            `json:"status"`
	ItemName      string          `json:"itemName"`
	ItemType      models.ItemType `json:"itemType"`
	ItemLinkedUrl string          `json:"itemLinkedUrl"`
	ItemResource  *string         `json:"itemResource"`
	IsGeneric     bool            `json:"isGeneric"`
	ProductSet    *[]int          `json:"productSet"`
}
