package services

import (
	"math"
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (serviceManager *ServiceManager) GetBuyerTypes(w http.ResponseWriter, r *http.Request) {
	var buyerTypes []models.BuyerType

	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	_ = serviceManager.db.Model(&buyerTypes).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&buyerTypes)
	response["buyerTypes"] = buyerTypes
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	response["totalItems"] = totalItems
	helpers.WriteJSON(w, http.StatusOK, response)
}
