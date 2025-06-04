package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

type VAUpdate struct {
	AlertSound string
	AlertColor string
}

type VABPost struct {
	AlertSound     string
	AlertColor     string
	BuyerTypeIdSet []int
}

type VAPPost struct {
	AlertSound     string
	AlertColor     string
	PromotionIdSet []int
}

func (sm *ServiceManager) GetVAB(w http.ResponseWriter, r *http.Request) {
	var vab []models.VerifierAlertBuyerType
	response := helpers.PaginateRequest(r, vab, sm.db, "buyerTypes")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetVABById(w http.ResponseWriter, r *http.Request) {
	vab := models.VerifierAlertBuyerType{}
	err := helpers.GetById(&vab, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no verifier alert with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, vab)
	}
}

func (sm *ServiceManager) GetVAP(w http.ResponseWriter, r *http.Request) {
	var vap []models.VerifierAlertPromotion
	response := helpers.PaginateRequest(r, vap, sm.db, "promotions")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetVAPById(w http.ResponseWriter, r *http.Request) {
	vap := models.VerifierAlertPromotion{}
	err := helpers.GetById(&vap, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no verifier alert with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, vap)
	}
}

func (sm *ServiceManager) PutVAB(w http.ResponseWriter, r *http.Request) {
	req := &VAUpdate{}
	vab := models.VerifierAlertBuyerType{}
	err := helpers.GetById(&vab, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no verifier alert with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	vabUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&vab).Updates(vabUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, vabUpdate)
}

func (sm *ServiceManager) PutVAP(w http.ResponseWriter, r *http.Request) {
	req := &VAUpdate{}
	vap := models.VerifierAlertPromotion{}
	err := helpers.GetById(&vap, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no verifier alert with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	vapUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&vap).Updates(vapUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, vapUpdate)
}

func (sm *ServiceManager) DeleteVAB(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := sm.db.Delete(&models.VerifierAlertBuyerType{}, id).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't delete"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{"success": "deleted properly"})
}

func (sm *ServiceManager) DeleteVAP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := sm.db.Delete(&models.VerifierAlertPromotion{}, id).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't delete"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{"success": "deleted properly"})
}

func (sm *ServiceManager) PostVAB(w http.ResponseWriter, r *http.Request) {
	req := &VABPost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	fmt.Println(req)
	for _, buyerTypeId := range req.BuyerTypeIdSet {
		vab := models.VerifierAlertBuyerType{
			BuyerTypeID: buyerTypeId,
			AlertColor:  req.AlertColor,
			AlertSound:  req.AlertSound,
		}
		err := sm.db.Create(&vab).Error
		fmt.Println("Test")
		if err != nil {
			helpers.WriteJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	// helpers.WriteJSON(w, http.StatusOK, map[string]string{"error": "Created verifier alerts properly"})
}

func (sm *ServiceManager) PostVAP(w http.ResponseWriter, r *http.Request) {
	req := &VAPPost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, promotionId := range req.PromotionIdSet {
		vab := models.VerifierAlertPromotion{
			PromotionID: promotionId,
			AlertColor:  req.AlertColor,
			AlertSound:  req.AlertSound,
		}
		err := sm.db.Create(&vab).Error
		if err != nil {
			helpers.WriteJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"error": "Created verifier alerts properly"})
}
