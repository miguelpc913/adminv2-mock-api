package services

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/go-chi/chi"
	dtoAffiliateAgreement "github.com/tiqueteo/adminv2-mock-api/api/dto/affiliateAgreement"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (serviceManager *ServiceManager) GetAffiliateAgreement(w http.ResponseWriter, r *http.Request) {
	var affiliateAgreements []models.AffiliateAgreement

	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	_ = serviceManager.db.Model(&affiliateAgreements).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&affiliateAgreements)
	response["affiliateAgreements"] = affiliateAgreements
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	response["totalItems"] = totalItems
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetAffiliateAgreementById(w http.ResponseWriter, r *http.Request) {
	var affiliateAgreement models.AffiliateAgreement
	id := chi.URLParam(r, "id")
	err := sm.db.Preload(clause.Associations).Find(&affiliateAgreement, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliateAgreement with that id"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
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
	var affiliateAgreement models.AffiliateAgreement
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&affiliateAgreement, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate item with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	affiliateAgreementUpdate := models.AffiliateAgreement{
		Status:          req.Status,
		Description:     req.Description,
		AgreementType:   req.AgreementType,
		ValueType:       req.ValueType,
		BaseReturnValue: req.BaseReturnValue,
		IsDefault:       req.IsDefault,
	}
	err = sm.db.Model(&affiliateAgreement).Where("affiliate_agreement_id = ?", id).Select("Status", "Description", "AgreementType", "ValueType", "BaseReturnValue", "IsDefault").Updates(affiliateAgreementUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}

func (sm *ServiceManager) PutAffiliateAgreementValities(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateAgreement.Validities{}
	var affiliateAgreement models.AffiliateAgreement
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&affiliateAgreement, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate item with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	affiliateAgreementUpdate := models.AffiliateAgreement{
		WeekDay:       req.WeekDay,
		DisabledDates: req.DisabledDates,
	}
	for _, disabledDate := range req.DisabledDates {
		_, err := helpers.ParseDate(disabledDate)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Disabled dates are not valid"})
			return
		}
	}
	err = sm.db.Model(&affiliateAgreement).Where("affiliate_agreement_id = ?", id).Updates(affiliateAgreementUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}

func (sm *ServiceManager) PutAffiliateAgreementBuyerTypes(w http.ResponseWriter, r *http.Request) {
	var affiliateAgreement models.AffiliateAgreement
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&affiliateAgreement, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate agreement with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	buyerTypes := []models.BuyerType{}
	for _, id := range req {
		buyerType := models.BuyerType{}
		if err := sm.db.First(&buyerType, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "BuyerTypes are not valid"})
			return
		}
		buyerTypes = append(buyerTypes, buyerType)
	}

	sm.db.Model(&affiliateAgreement).Association("BuyerTypeSet").Replace(buyerTypes)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutAffiliateAgreementProducts(w http.ResponseWriter, r *http.Request) {
	var affiliateAgreement models.AffiliateAgreement
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&affiliateAgreement, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no affiliate agreement with that id"})
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

	sm.db.Model(&affiliateAgreement).Association("ProductSet").Replace(products)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PostAffiliateAgreement(w http.ResponseWriter, r *http.Request) {
	req := &dtoAffiliateAgreement.Post{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	//Manage associations
	products := []models.Product{}
	for _, productId := range req.ProductSet {
		product := models.Product{}
		if err := sm.db.First(&product, productId).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "products are not valid"})
			return
		}
		products = append(products, product)
	}
	buyerTypes := []models.BuyerType{}
	for _, buyerTypeId := range req.BuyerTypeSet {
		buyerType := models.BuyerType{}
		if err := sm.db.First(&buyerType, buyerTypeId).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "buyerTypes are not valid"})
			return
		}
		buyerTypes = append(buyerTypes, buyerType)
	}
	for _, disabledDate := range req.DisabledDates {
		_, err := helpers.ParseDate(disabledDate)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Disabled dates are not valid"})
			return
		}
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
	err := sm.db.Create(&affiliateAgreement).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}
