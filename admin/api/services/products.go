package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	response := helpers.PaginateRequest(r, products, sm.db, "products")
	helpers.WriteJSON(w, http.StatusOK, response)
}
