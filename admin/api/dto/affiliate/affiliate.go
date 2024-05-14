package dtoAffiliate

import "github.com/tiqueteo/adminv2-mock-api/db/models"

type Status struct {
	Status models.AffiliateStatus `json:"status"`
}
