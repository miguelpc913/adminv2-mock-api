package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	dtoBuyerTypeRules "github.com/tiqueteo/adminv2-mock-api/api/dto/buyerTypeRules"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (sm *ServiceManager) GetBuyerTypesRules(w http.ResponseWriter, r *http.Request) {
	var buyerTypeRules []models.BuyerTypeRule
	response := helpers.PaginateRequest(r, buyerTypeRules, sm.db, "buyerTypeRules")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetBuyerTypeRulesId(w http.ResponseWriter, r *http.Request) {
	buyerTypeRule := models.BuyerTypeRule{}
	err := helpers.GetById(&buyerTypeRule, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no entity with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, buyerTypeRule)
	}
}

func (sm *ServiceManager) PutBuyerTypeRulesIdentity(w http.ResponseWriter, r *http.Request) {
	req := dtoBuyerTypeRules.BuyerTypeRulesIdentity{}
	buyerTypeRule := models.BuyerTypeRule{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	if err := helpers.GetById(&buyerTypeRule, r, sm.db); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "No buyer type rule identity with that id"})
		return
	}
	buyerTypeRuleUpdate := helpers.StructToMap(req)
	if err := sm.db.Model(&buyerTypeRule).Updates(buyerTypeRuleUpdate).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, buyerTypeRule)
}
func (sm *ServiceManager) PutBuyerTypeRuleConfiguration(w http.ResponseWriter, r *http.Request) {
	req := dtoBuyerTypeRules.BuyerTypeRuleConfiguration{}
	buyerTypeRule := models.BuyerTypeRule{}
	vars := models.Vars{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	if err := helpers.GetById(&buyerTypeRule, r, sm.db); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "No buyer type rule identity with that id"})
		return
	}
	if err := sm.db.Preload(clause.Associations).First(&vars, buyerTypeRule.Vars.VarId).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "No buyer type rule identity with that vars id"})
		return
	}
	if err := sm.db.Model(&buyerTypeRule).Updates(map[string]interface{}{"ErrorMessage": req.ErrorMessage, "Priority": req.Priority}).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update buyer type rule"})
		return
	}
	varsData := models.Vars{
		VarId:           vars.VarId,
		X:               req.Vars.X,
		Y:               req.Vars.Y,
		M:               req.Vars.M,
		N:               req.Vars.N,
		BuyerTypeRuleID: buyerTypeRule.BuyerTypeRuleID,
	}
	if err := sm.db.Model(&vars).Updates(varsData).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, buyerTypeRule)
}

func (sm *ServiceManager) PutBuyerTypeRulesProductSet(w http.ResponseWriter, r *http.Request) {
	var buyerTypeRule models.BuyerTypeRule
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&buyerTypeRule, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not buyer type rule with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	products := []models.Product{}
	for _, id := range req {
		product := models.Product{}
		if err := sm.db.First(&product, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Products are not valid"})
			return
		}
		products = append(products, product)
	}

	sm.db.Model(&buyerTypeRule).Association("ProductSet").Replace(products)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PostBuyerTypeRules(w http.ResponseWriter, r *http.Request) {
	req := dtoBuyerTypeRules.BuyerTypeRulePostDTO{}
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
	buyerTypeRule := models.BuyerTypeRule{
		Status:                  req.Status,
		Name:                    req.Name,
		BuyerTypeRuleTemplateID: req.BuyerTypeRuleTemplateID,
		ErrorMessage:            req.ErrorMessage,
		Priority:                req.Priority,
		ProductSet:              products,
	}
	if err := sm.db.Create(&buyerTypeRule).Error; err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	varsCopy := models.Vars{
		X: req.Vars.X,
		Y: req.Vars.Y,
		M: req.Vars.M,
		N: req.Vars.N,
	}

	varsCopy.BuyerTypeRuleID = buyerTypeRule.BuyerTypeRuleID
	if err := sm.db.Create(&varsCopy).Error; err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, buyerTypeRule)
}
