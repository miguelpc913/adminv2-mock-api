package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetBuyerTypes(w http.ResponseWriter, r *http.Request) {
	var buyerTypes []models.BuyerType
	response := helpers.PaginateRequest(r, buyerTypes, sm.db, "buyerTypes")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetBuyerTypeById(w http.ResponseWriter, r *http.Request) {
	buyerType := models.BuyerType{}
	err := helpers.GetById(&buyerType, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no buyerType with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, buyerType)
	}
}
