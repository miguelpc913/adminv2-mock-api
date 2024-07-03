package services

import (
	"encoding/json"
	"net/http"

	dtoRr "github.com/tiqueteo/adminv2-mock-api/api/dto/recommendationRules"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetRecommendationRules(w http.ResponseWriter, r *http.Request) {
	var recommendations []models.RecommendationRule
	response := helpers.PaginateRequest(r, recommendations, sm.db, "recommendationRules")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetRecommendationRuleById(w http.ResponseWriter, r *http.Request) {
	var recommendationRule models.RecommendationRule
	err := helpers.GetById(&recommendationRule, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no recommendation rule with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, recommendationRule)
	}
}

func (sm *ServiceManager) PutOrderRecommendationRules(w http.ResponseWriter, r *http.Request) {
	var recommendationRules []models.RecommendationRule
	var req []dtoRr.DisplayOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, item := range req {
		err := sm.db.Model(&recommendationRules).Where("recommendation_rule_id = ?", item.RecommendationRuleId).Update("priority", item.Priority).Error
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid recommendation id"})
		}
	}

	helpers.WriteJSON(w, http.StatusOK, recommendationRules)
}

func (sm *ServiceManager) PutRecommendationRuleIdentity(w http.ResponseWriter, r *http.Request) {
	var req dtoRr.PutRecommendationIdentity
	var recommendationRule models.RecommendationRule
	err := helpers.GetById(&recommendationRule, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no recommendation rule with that id"})
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	recommendationRuleUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&recommendationRule).Updates(recommendationRuleUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, recommendationRule)
}

func (sm *ServiceManager) PutRecommendationRuleGeneral(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationGeneral
	err := helpers.GetById(&recommendation, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no recommendation rule with that id"})
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	recommendationRuleUpdate := models.RecommendationRule{
		ProductId:        req.ProductId,
		OfferedProductId: req.OfferedProductId,
		DirectAddToCart:  req.DirectAddToCart,
	}
	err = sm.db.Model(&recommendation).Select("ProductId", "OfferedProductId", "DirectAddToCart").Updates(recommendationRuleUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationRuleValidities(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationValidities
	err := helpers.GetById(&recommendation, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no recommendation rule with that id"})
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
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
	_, err = helpers.ParseTime(req.StartTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartTime is not valid"})
		return
	}
	_, err = helpers.ParseTime(req.EndTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime is not valid"})
		return
	}
	recommendationRuleUpdate := models.RecommendationRule{
		StartDatetime:              startDateTime,
		EndDatetime:                endDatetime,
		EventStartDatetime:         eventStartDatetime,
		EventEndDatetime:           eventEndDatetime,
		WeekDay:                    req.WeekDay,
		StartTime:                  req.StartTime,
		EndTime:                    req.EndTime,
		SessionOffsetMinutesBefore: req.SessionOffsetMinutesBefore,
		SessionOffsetMinutesAfter:  req.SessionOffsetMinutesAfter,
	}
	err = sm.db.Model(&recommendation).Updates(recommendationRuleUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationDisplay(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	var req dtoRr.PutRecommendationDisplay
	err := helpers.GetById(&recommendation, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not recommendation with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	recommendationRuleUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&recommendation).Updates(recommendationRuleUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, recommendation)
}

func (sm *ServiceManager) PutRecommendationSalesGroups(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	err := helpers.UpdateRelation(r, recommendation, models.SalesGroup{}, sm.db, "SalesGroupSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PutRecommendationBuyerTypes(w http.ResponseWriter, r *http.Request) {
	var recommendation models.RecommendationRule
	err := helpers.UpdateRelation(r, recommendation, models.BuyerType{}, sm.db, "BuyerTypeSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PostRecommendationRule(w http.ResponseWriter, r *http.Request) {
	req := &dtoRr.PostRecommendation{}
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
	_, err = helpers.ParseTime(req.StartTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "StartTime is not valid"})
		return
	}
	_, err = helpers.ParseTime(req.EndTime)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime is not valid"})
		return
	}

	//Manage associations
	salesGroups := []models.SalesGroup{}
	err = helpers.GetByIds(&salesGroups, req.SalesGroupSet, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "salesgroups are not valid"})
		return
	}
	buyerTypes := []models.BuyerType{}
	err = helpers.GetByIds(&buyerTypes, req.BuyerTypeSet, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "buyerTypes are not valid"})
		return
	}

	recommendationRule := models.RecommendationRule{
		Status:                     req.Status,
		Name:                       req.Name,
		OfferingType:               req.OfferingType,
		ProductId:                  req.ProductId,
		OfferedProductId:           req.OfferedProductId,
		DirectAddToCart:            req.DirectAddToCart,
		StartDatetime:              startDateTime,
		EndDatetime:                endDatetime,
		EventStartDatetime:         eventStartDatetime,
		EventEndDatetime:           eventEndDatetime,
		WeekDay:                    req.WeekDay,
		StartTime:                  req.StartTime,
		EndTime:                    req.EndTime,
		SessionOffsetMinutesBefore: req.SessionOffsetMinutesBefore,
		SessionOffsetMinutesAfter:  req.SessionOffsetMinutesAfter,
		SalesGroupSet:              salesGroups,
		BuyerTypeSet:               buyerTypes,
		Title:                      req.Title,
		Body:                       req.Body,
		Footer:                     req.Footer,
	}
	err = sm.db.Create(&recommendationRule).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, recommendationRule)
}
