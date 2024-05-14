package services

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/go-chi/chi"
	dtoAffiliate "github.com/tiqueteo/adminv2-mock-api/api/dto/affiliate"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (sm *ServiceManager) GetAffiliates(w http.ResponseWriter, r *http.Request) {
	var affiliates []models.Affiliate

	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	status := r.URL.Query().Get("status")
	if status != "" {
		_ = sm.db.Preload(clause.Associations).Model(&affiliates).Where("status = ?", status).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&affiliates)
	} else {
		_ = sm.db.Preload(clause.Associations).Model(&affiliates).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&affiliates)
	}
	response["affiliates"] = affiliates
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	response["totalItems"] = totalItems
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetAffiliateById(w http.ResponseWriter, r *http.Request) {
	affiliate := models.Affiliate{}
	id := chi.URLParam(r, "id")

	err := sm.db.Preload(clause.Associations).Find(&affiliate, id).Error

	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliateAgreement with that id"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliate)
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
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&affiliate, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	affiliateAgreementSet := []models.AffiliateAgreement{}
	for _, id := range req {
		agreement := models.AffiliateAgreement{}
		if err := sm.db.First(&agreement, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Affiliate agreements are not valid"})
			return
		}
		affiliateAgreementSet = append(affiliateAgreementSet, agreement)
	}

	sm.db.Model(&affiliate).Association("AffiliateAgreementSet").Replace(affiliateAgreementSet)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}
