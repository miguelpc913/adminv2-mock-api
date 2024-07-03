package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	dtoAffiliateItem "github.com/tiqueteo/adminv2-mock-api/api/dto/affiliateItems"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (sm *ServiceManager) GetAffiliateItems(w http.ResponseWriter, r *http.Request) {
	var affiliateItems []models.AffiliateItem
	response := helpers.PaginateRequest(r, affiliateItems, sm.db, "affiliateItems")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetAffiliateItemById(w http.ResponseWriter, r *http.Request) {
	var affiliateItem models.AffiliateItem
	id := chi.URLParam(r, "id")
	err := sm.db.Preload(clause.Associations).Find(&affiliateItem, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliateItem with that id"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateItem)
}

func (sm *ServiceManager) PostAffiliateItem(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateItem.AffiliateItemReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	products := []models.Product{}
	if req.ProductSet != nil && len(*req.ProductSet) > 0 {
		err := helpers.GetByIds(&products, *req.ProductSet, sm.db)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "products are not valid"})
			return
		}
	}
	affiliateItem := models.AffiliateItem{
		Status:        req.Status,
		ItemName:      req.ItemName,
		ItemType:      req.ItemType,
		ItemLinkedUrl: req.ItemLinkedUrl,
		ItemResource:  req.ItemResource,
		IsGeneric:     req.IsGeneric,
		ProductSet:    products,
	}
	err := sm.db.Create(&affiliateItem).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateItem)
}

func (sm *ServiceManager) PutAffiliateItem(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateItem.AffiliateItemReq{}
	var affiliateItem models.AffiliateItem
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&affiliateItem, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate item with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	products := []models.Product{}
	if req.ProductSet != nil && len(*req.ProductSet) > 0 {
		err := helpers.GetByIds(&products, *req.ProductSet, sm.db)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "products are not valid"})
			return
		}
		sm.db.Model(&affiliateItem).Association("ProductSet").Replace(products)
	}
	affiliateItemUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&affiliateItem).Updates(affiliateItemUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateItem)
}
