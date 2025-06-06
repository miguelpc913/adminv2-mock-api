package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetVenueCapacities(w http.ResponseWriter, r *http.Request) {
	var venues []models.VenueCapacity
	response := helpers.PaginateRequest(r, venues, sm.db, "venueCapacities")
	helpers.WriteJSON(w, http.StatusOK, response)
}
