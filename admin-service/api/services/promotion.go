package services

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	dtoPromotion "github.com/tiqueteo/adminv2-mock-api/api/dto/promotion"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (sm *ServiceManager) GetPromotions(w http.ResponseWriter, r *http.Request) {
	var promotions []models.Promotion
	pagination := helpers.GeneratePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	sm.db.Preload(clause.Associations).Model(&promotions).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&promotions).Preload(clause.Associations)
	response["promotions"] = promotions
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetPromotionById(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	id := chi.URLParam(r, "id")
	err := sm.db.Preload(clause.Associations).Find(&promotion, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no promotion with that id"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, promotion)
}

func (sm *ServiceManager) PutPromotionIdentity(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	var req dtoPromotion.PutPromotionIdentity
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Model(promotion).Where("promotion_id = ?", id).Select("Name", "Status", "ShortName").Updates(req).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, promotion)
}

func (sm *ServiceManager) PutPromotionGeneral(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	var req dtoPromotion.PutPromotionGeneral
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	var promotionCodes []models.PromotionalCode
	promotion_id, err := strconv.Atoi(id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Incorrect promotion id"})
		return
	}
	err = sm.db.Find(&promotion, promotion_id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no promotion with that id"})
		return
	}
	if promotion.CodeType == "generated" {
		for i := 0; i < req.NumberOfCodes; i++ {
			generatedCode := helpers.RandStringBytes(promotion.CodeLength)
			promotionCode := models.PromotionalCode{
				Code:        generatedCode,
				Quantity:    strconv.Itoa(promotion.Quantity),
				PromotionId: promotion.PromotionId,
			}
			promotionCodes = append(promotionCodes, promotionCode)
		}
	} else {
		for _, promotionCode := range req.PromotionalCodeSet {
			promotionCode := models.PromotionalCode{
				Code:        promotionCode.Code,
				Quantity:    promotionCode.Quantity,
				PromotionId: promotion_id,
			}
			promotionCodes = append(promotionCodes, promotionCode)
		}
	}
	err = sm.db.Create(&promotionCodes).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	sm.db.Model(&promotion).Association("PromotionalCodeSet").Replace(promotionCodes)
	err = sm.db.Model(&promotion).Where("promotion_id = ?", promotion_id).Select("Amount", "Percentage", "LeftPurchased", "RightPaid", "IsPromotionAffiliateEnabled", "HideAmountAtTicket", "ShowOriginalAmountAtTicket", "IsGrouped", "RedeemType", "CodeType", "NumberOfCodes", "CodeLength", "Quantity").Updates(models.Promotion{Amount: req.Amount, Percentage: req.Percentage, LeftPurchased: req.LeftPurchased, RightPaid: req.RightPaid, IsPromotionAffiliateEnabled: req.IsPromotionAffiliateEnabled, HideAmountAtTicket: req.HideAmountAtTicket, ShowOriginalAmountAtTicket: req.ShowOriginalAmountAtTicket, IsGrouped: req.IsGrouped, RedeemType: req.RedeemType, CodeType: req.CodeType, NumberOfCodes: req.NumberOfCodes, CodeLength: req.CodeLength, Quantity: req.Quantity}).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, promotion)
}

func (sm *ServiceManager) PutPromotionValidities(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	var req dtoPromotion.PutPromotionValidities
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Find(&promotion, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not promotion with that id"})
		return
	}
	startDateTime, err := helpers.ParseDateTime(req.StartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartDatetime is not valid"})
		return
	}
	endDatetime, err := helpers.ParseDateTime(req.EndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndDatetime is not valid"})
		return
	}
	eventStartDatetime, err := helpers.ParseDateTime(req.EventStartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventStartDatetime is not valid"})
		return
	}
	eventEndDatetime, err := helpers.ParseDateTime(req.EventEndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventEndDatetime is not valid"})
		return
	}

	promotion = models.Promotion{
		StartDatetime:         startDateTime,
		EndDatetime:           endDatetime,
		EventStartDatetime:    eventStartDatetime,
		EventEndDatetime:      eventEndDatetime,
		WeekDay:               req.WeekDay,
		MinSecondsBeforeEvent: req.MinSecondsBeforeEvent,
		MaxSecondsBeforeEvent: req.MaxSecondsBeforeEvent,
		DisabledDates:         req.DisabledDates,
	}
	err = sm.db.Model(promotion).Where("promotion_id = ?", id).Updates(promotion).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, promotion)
}

func (sm *ServiceManager) PutPromotionAdvancedSettings(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	var req dtoPromotion.PutPromotionAdvancedSettings
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	promotion_id, err := strconv.Atoi(id)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Incorrect promotion id"})
		return
	}
	sm.db.First(&promotion, promotion_id)
	var promotionPrices []models.PromotionPrice
	for _, promotionPrice := range req.PromotionPriceSet {
		newPromotionPrice := models.PromotionPrice{
			PromotionId:  promotion_id,
			BuyerTypeId:  promotionPrice.BuyerTypeID,
			SalesGroupId: promotionPrice.SalesGroupId,
			Amount:       promotionPrice.Amount,
			Percentage:   promotionPrice.Percentage,
		}
		promotionPrices = append(promotionPrices, newPromotionPrice)
	}
	err = sm.db.Create(&promotionPrices).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	sm.db.Model(&promotion).Association("PromotionPriceSet").Replace(promotionPrices)
	helpers.WriteJSON(w, http.StatusOK, promotion)
}

func (sm *ServiceManager) PutPromotionSalesGroups(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&promotion, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no promotion info with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	salesGroups := []models.SalesGroup{}
	for _, id := range req {
		salesGroup := models.SalesGroup{}
		if err := sm.db.First(&salesGroup, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "SalesGroups are not valid"})
			return
		}
		salesGroups = append(salesGroups, salesGroup)
	}

	sm.db.Model(&promotion).Association("SalesGroupSet").Replace(salesGroups)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutPromotionBuyerTypes(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&promotion, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no promotion info with that id"})
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

	sm.db.Model(&promotion).Association("BuyerTypeSet").Replace(buyerTypes)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutPromotionProducts(w http.ResponseWriter, r *http.Request) {
	var promotion models.Promotion
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&promotion, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no promotion info with that id"})
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

	sm.db.Model(&promotion).Association("ProductSet").Replace(products)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PostValidateCode(w http.ResponseWriter, r *http.Request) {
	req := &dtoPromotion.CodeToValidate{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	promotionalCode := &models.PromotionalCode{}

	err := sm.db.Model(promotionalCode).Where("code = ?", req.Code).First(&promotionalCode).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		helpers.WriteJSON(w, http.StatusOK, "code is valid")
	} else {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid code"})
	}
}

func (sm *ServiceManager) PostPromotion(w http.ResponseWriter, r *http.Request) {
	req := &dtoPromotion.PromotionPost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	//Validate dates
	startDateTime, err := helpers.ParseDateTime(req.StartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartDatetime is not valid"})
		return
	}
	endDatetime, err := helpers.ParseDateTime(req.EndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndDatetime is not valid"})
		return
	}
	eventStartDatetime, err := helpers.ParseDateTime(req.EventStartDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventStartDatetime is not valid"})
		return
	}
	eventEndDatetime, err := helpers.ParseDateTime(req.EventEndDatetime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EventEndDatetime is not valid"})
		return
	}

	//Manage associations
	salesGroups := []models.SalesGroup{}
	for _, saleGroupData := range req.SalesGroupSet {
		salesGroup := models.SalesGroup{}
		if err := sm.db.First(&salesGroup, saleGroupData.SalesGroupId).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "SalesGroups are not valid"})
			return
		}
		salesGroups = append(salesGroups, salesGroup)
	}

	buyerTypes := []models.BuyerType{}
	for _, buyerTypeData := range req.BuyerTypeSet {
		buyerType := models.BuyerType{}
		if err := sm.db.First(&buyerType, buyerTypeData.BuyerTypeID).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "buyerTypes are not valid"})
			return
		}
		buyerTypes = append(buyerTypes, buyerType)
	}

	products := []models.Product{}
	for _, productData := range req.ProductSet {
		product := models.Product{}
		if err := sm.db.First(&product, productData.ProductId).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "buyerTypes are not valid"})
			return
		}
		products = append(products, product)
	}
	for _, disabledDate := range req.DisabledDates {
		_, err := helpers.ParseDate(disabledDate)
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Selected dates is not valid"})
			return
		}
	}
	promotion := models.Promotion{
		Status:                      req.Status,
		Name:                        req.Name,
		ShortName:                   req.ShortName,
		PromotionType:               req.PromotionType,
		Amount:                      req.Amount,
		Percentage:                  req.Percentage,
		LeftPurchased:               req.LeftPurchased,
		RightPaid:                   req.RightPaid,
		IsPromotionAffiliateEnabled: req.IsPromotionAffiliateEnabled,
		HideAmountAtTicket:          req.HideAmountAtTicket,
		ShowOriginalAmountAtTicket:  req.ShowOriginalAmountAtTicket,
		IsGrouped:                   req.IsGrouped,
		RedeemType:                  req.RedeemType,
		CodeType:                    req.CodeType,
		NumberOfCodes:               req.NumberOfCodes,
		CodeLength:                  req.CodeLength,
		Quantity:                    req.Quantity,
		StartDatetime:               startDateTime,
		EndDatetime:                 endDatetime,
		MinSecondsBeforeEvent:       req.MinSecondsBeforeEvent,
		MaxSecondsBeforeEvent:       req.MaxSecondsBeforeEvent,
		EventStartDatetime:          eventStartDatetime,
		EventEndDatetime:            eventEndDatetime,
		WeekDay:                     req.WeekDay,
		DisabledDates:               req.DisabledDates,
		SalesGroupSet:               salesGroups,
		BuyerTypeSet:                buyerTypes,
		ProductSet:                  products,
	}
	err = sm.db.Create(&promotion).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	if promotion.CodeType == "generated" {
		var promotionalCodes []models.PromotionalCode
		for i := 0; i < promotion.NumberOfCodes; i++ {
			generatedCode := helpers.RandStringBytes(promotion.CodeLength)
			promotionCode := models.PromotionalCode{
				Code:        generatedCode,
				Quantity:    strconv.Itoa(promotion.Quantity),
				PromotionId: promotion.PromotionId,
			}
			promotionalCodes = append(promotionalCodes, promotionCode)
		}
		err = sm.db.Create(&promotionalCodes).Error
		if err != nil {
			helpers.WriteJSON(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		for _, promotionCode := range req.PromotionalCodeSet {
			promotionCode := models.PromotionalCode{
				Code:        promotionCode.Code,
				Quantity:    promotionCode.Quantity,
				PromotionId: promotion.PromotionId,
			}
			err = sm.db.Create(&promotionCode).Error
			if err != nil {
				helpers.WriteJSON(w, http.StatusInternalServerError, err)
				return
			}
		}
	}

	helpers.WriteJSON(w, http.StatusOK, promotion)
}
