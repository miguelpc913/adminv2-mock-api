package services

import (
	"encoding/json"
	"net/http"

	dtoAffiliateAgreement "github.com/tiqueteo/adminv2-mock-api/api/dto/affiliateAgreement"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetAffiliateAgreement(w http.ResponseWriter, r *http.Request) {
	var affiliateAgreements []models.AffiliateAgreement
	response := helpers.PaginateRequest(r, affiliateAgreements, sm.db, "affiliateAgreements")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetAffiliateAgreementById(w http.ResponseWriter, r *http.Request) {
	AffiliateAgreement := models.AffiliateAgreement{}
	err := helpers.GetById(&AffiliateAgreement, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliateAgreement with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, AffiliateAgreement)
	}
}

func (sm *ServiceManager) PutOrderAffiliateAgreement(w http.ResponseWriter, r *http.Request) {
	var affiliateAgreements []models.AffiliateAgreement
	var req []dtoAffiliateAgreement.DisplayOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, item := range req {
		err := sm.db.Model(&affiliateAgreements).Where("affiliate_agreement_id = ?", item.AffiliateAgreementId).Update("priority", item.Priority).Error
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid affiliate agreement id"})
		}
	}

	helpers.WriteJSON(w, http.StatusOK, affiliateAgreements)
}

func (sm *ServiceManager) PutAffiliateAgreementGeneral(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateAgreement.General{}
	affiliateAgreement := models.AffiliateAgreement{}
	err := helpers.GetById(&affiliateAgreement, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate item with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	affiliateAgreementUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&affiliateAgreement).Updates(affiliateAgreementUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}

func (sm *ServiceManager) PutAffiliateAgreementValities(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateAgreement.Validities{}
	affiliateAgreement := models.AffiliateAgreement{}
	err := helpers.GetById(&affiliateAgreement, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate item with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err = helpers.ValidateDates(req.DisabledDates)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Disabled dates are not valid"})
		return
	}
	affiliateAgreementUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&affiliateAgreement).Updates(affiliateAgreementUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}

func (sm *ServiceManager) PutAffiliateAgreementBuyerTypes(w http.ResponseWriter, r *http.Request) {
	affiliateAgreement := models.AffiliateAgreement{}
	err := helpers.UpdateRelation(r, affiliateAgreement, models.BuyerType{}, sm.db, "BuyerTypeSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PutAffiliateAgreementProducts(w http.ResponseWriter, r *http.Request) {
	affiliateAgreement := models.AffiliateAgreement{}
	err := helpers.UpdateRelation(r, affiliateAgreement, models.Product{}, sm.db, "ProductSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PostAffiliateAgreement(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateAgreement.Post{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	products := []models.Product{}
	err := helpers.GetByIds(&products, req.ProductSet, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "products are not valid"})
		return
	}
	buyerTypes := []models.BuyerType{}
	err = helpers.GetByIds(&buyerTypes, req.BuyerTypeSet, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "buyerTypes are not valid"})
		return
	}
	err = helpers.ValidateDates(req.DisabledDates)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Disabled dates are not valid"})
		return
	}
	affiliateAgreement := models.AffiliateAgreement{
		Status:          req.Status,
		Description:     req.Description,
		AgreementType:   req.AgreementType,
		ValueType:       req.ValueType,
		BaseReturnValue: req.BaseReturnValue,
		WeekDay:         req.WeekDay,
		IsDefault:       req.IsDefault,
		DisabledDates:   req.DisabledDates,
		ProductSet:      products,
		BuyerTypeSet:    buyerTypes,
	}
	err = sm.db.Create(&affiliateAgreement).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}
