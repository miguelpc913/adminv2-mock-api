package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetExtras(w http.ResponseWriter, r *http.Request) {
	var extras []models.Extra
	response := helpers.PaginateRequest(r, extras, sm.db, "extras")
	helpers.WriteJSON(w, http.StatusOK, response)
}
