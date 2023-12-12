package services

import (
	"math"
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (serviceManager *ServiceManager) GetProductTags(w http.ResponseWriter, r *http.Request) {
	var productTags []models.ProductTag
	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	_ = serviceManager.db.Model(&productTags).Preload("SalesGroupSet").Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&productTags).Error
	response["productTags"] = productTags
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (serviceManager *ServiceManager) PutOrderProductTags(w http.ResponseWriter, r *http.Request) {
	var productTags []models.ProductTag
	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	_ = serviceManager.db.Model(&productTags).Preload("SalesGroupSet").Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&productTags).Error
	response["productTags"] = productTags
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	helpers.WriteJSON(w, http.StatusOK, response)
}
