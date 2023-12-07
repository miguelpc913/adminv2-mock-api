package services

import (
	"admin-v2/api/helpers"
	"admin-v2/db/models"
	"math"
	"net/http"
)

func (serviceManager *ServiceManager) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	_ = serviceManager.db.Model(&products).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&products)
	response["products"] = products
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	response["totalItems"] = totalItems
	helpers.WriteJSON(w, http.StatusOK, response)
}
