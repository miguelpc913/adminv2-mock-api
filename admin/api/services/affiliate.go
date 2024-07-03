package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	dtoAffiliate "github.com/tiqueteo/adminv2-mock-api/api/dto/affiliate"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetAffiliates(w http.ResponseWriter, r *http.Request) {
	var affiliates []models.Affiliate
	response := helpers.PaginateRequest(r, affiliates, sm.db, "affiliates")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetAffiliateById(w http.ResponseWriter, r *http.Request) {
	affiliate := models.Affiliate{}
	err := helpers.GetById(&affiliate, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no entity with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, affiliate)
	}
}

func (sm *ServiceManager) PutAffiliate(w http.ResponseWriter, r *http.Request) {
	affiliate := models.Affiliate{}
	req := dtoAffiliate.Status{}
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Model(affiliate).Where("affiliate_id = ?", id).Select("Status").Updates(req).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, affiliate)
}

func (sm *ServiceManager) PutAffiliateAgreements(w http.ResponseWriter, r *http.Request) {
	var affiliate models.Affiliate
	err := helpers.UpdateRelation(r, affiliate, models.AffiliateAgreement{}, sm.db, "AffiliateAgreementSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}
