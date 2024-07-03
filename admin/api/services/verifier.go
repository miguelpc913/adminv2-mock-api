package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetVerifiers(w http.ResponseWriter, r *http.Request) {
	var verifiers []models.Verifier
	response := helpers.PaginateRequest(r, verifiers, sm.db, "verifiers")
	helpers.WriteJSON(w, http.StatusOK, response)
}
