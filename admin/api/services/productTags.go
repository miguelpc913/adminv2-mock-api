package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetProductTags(w http.ResponseWriter, r *http.Request) {
	var productTags []models.ProductTag
	response := helpers.PaginateRequest(r, productTags, sm.db, "productTags")
	helpers.WriteJSON(w, http.StatusOK, response)
}
