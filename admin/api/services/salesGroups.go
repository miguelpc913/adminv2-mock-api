package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetSalesGroups(w http.ResponseWriter, r *http.Request) {
	var salesGroups []models.SalesGroup
	response := helpers.PaginateRequest(r, salesGroups, sm.db, "salesGroups")
	helpers.WriteJSON(w, http.StatusOK, response)
}
