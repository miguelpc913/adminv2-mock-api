package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetReports(w http.ResponseWriter, r *http.Request) {
	var reports []models.Report
	response := helpers.PaginateRequest(r, reports, sm.db, "reports")
	helpers.WriteJSON(w, http.StatusOK, response)
}
