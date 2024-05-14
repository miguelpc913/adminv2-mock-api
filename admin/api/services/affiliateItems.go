package services

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/go-chi/chi"
	dtoAffiliateItem "github.com/tiqueteo/adminv2-mock-api/api/dto/affiliateItems"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (serviceManager *ServiceManager) GetAffiliateItems(w http.ResponseWriter, r *http.Request) {
	var affiliateItems []models.AffiliateItem

	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	_ = serviceManager.db.Model(&affiliateItems).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&affiliateItems)
	response["affiliateItems"] = affiliateItems
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	response["totalItems"] = totalItems
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
	//Manage associations
	if req.ProductSet != nil && len(*req.ProductSet) > 0 {
		for _, productId := range *req.ProductSet {
			product := models.Product{}
			if err := sm.db.First(&product, productId).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "products are not valid"})
				return
			}
			products = append(products, product)
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

type AffiliateItemReqProduct struct {
	ProductSet *[]int `json:"productSet"`
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
	if req.ProductSet != nil && len(*req.ProductSet) > 0 {
		//Manage associations
		products := []models.Product{}
		for _, productId := range *req.ProductSet {
			product := models.Product{}
			if err := sm.db.First(&product, productId).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "products are not valid"})
				return
			}
			products = append(products, product)
		}
		sm.db.Model(&affiliateItem).Association("ProductSet").Replace(products)
	}
	affiliateItemUpdate := models.AffiliateItem{
		Status:        req.Status,
		ItemName:      req.ItemName,
		ItemType:      req.ItemType,
		ItemLinkedUrl: req.ItemLinkedUrl,
		ItemResource:  req.ItemResource,
		IsGeneric:     req.IsGeneric,
	}
	err = sm.db.Model(&affiliateItem).Where("affiliate_item_id = ?", id).Select("Status", "ItemName", "ItemType", "ItemLinkedUrl", "ItemResource", "IsGeneric").Updates(affiliateItemUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateItem)
}
