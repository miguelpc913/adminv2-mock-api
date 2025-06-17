package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetPO(w http.ResponseWriter, r *http.Request) {
	var pointOfSales []models.PointsOfSale
	response := helpers.PaginateRequest(r, pointOfSales, sm.db, "pointOfSales")
	helpers.WriteJSON(w, http.StatusOK, response)
}
