package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetVenues(w http.ResponseWriter, r *http.Request) {
	var venues []models.Venue
	response := helpers.PaginateRequest(r, venues, sm.db, "venues")
	helpers.WriteJSON(w, http.StatusOK, response)
}
