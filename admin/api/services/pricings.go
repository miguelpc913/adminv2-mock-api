package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	dto "github.com/tiqueteo/adminv2-mock-api/api/dto/pricings"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (sm *ServiceManager) GetPricings(w http.ResponseWriter, r *http.Request) {
	pricings := []models.MainPricing{}
	productId := r.URL.Query().Get("productId")
	sm.db.Preload("Pricings.ProductExtraBuyerTypes").Preload("Pricings.DynamicPricingConfiguration").Preload("Pricings.DynamicPricingConfiguration.OccupancyRanges").Preload("Pricings.RecurrentTime").Preload("Pricings.ProductVenueBuyerTypes").Preload(clause.Associations).Model(&pricings).Where("product_id = ?", productId).Find(&pricings)
	helpers.WriteJSON(w, http.StatusOK, pricings)
}

func (sm *ServiceManager) GetPricingById(w http.ResponseWriter, r *http.Request) {
	mainPricing := models.MainPricing{}
	id := chi.URLParam(r, "id")
	err := sm.db.Preload("Pricings.ProductExtraBuyerTypes").Preload("Pricings.DynamicPricingConfiguration").Preload("Pricings.DynamicPricingConfiguration.OccupancyRanges").Preload("Pricings.RecurrentTime").Preload("Pricings.ProductVenueBuyerTypes").Preload(clause.Associations).Model(&mainPricing).First(&mainPricing, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no entity with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, mainPricing)
	}
}

func (sm *ServiceManager) PostBasePricing(w http.ResponseWriter, r *http.Request) {
	var req = dto.BasePricingPost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	_, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	_, err = time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	newBasePricing := models.MainPricing{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		ProductId: 2,
		Color:     req.Color,
		Name:      req.Name,
	}

	if err := sm.db.Create(&newBasePricing).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new base pricing"})
		return
	}
	prevMainPricing := models.MainPricing{}
	if err := sm.db.Model(&prevMainPricing).Preload(clause.Associations).Where("main_pricing_id = ?", newBasePricing.MainPricingId-1).Find(&prevMainPricing).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "can't find main pricing"})
		return
	}

	if err := createBaseSpecificPricing(newBasePricing, prevMainPricing, sm); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "can't create base specific pricing"})
		return
	}
}

func (sm *ServiceManager) PostSpecficPricing(w http.ResponseWriter, r *http.Request) {
	var req = dto.SpecificPricingDTO{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	mainPricingId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid main pricing id"})
		return
	}
	newSpecificPricing := models.SpecificPricing{
		MainPricingId: mainPricingId,
		Name:          req.Name,
		Weekdays:      req.WeekDays,
		EnabledDates:  req.EnabledDates,
		StartHour:     req.StartHour,
		EndHour:       req.EndHour,
		Default:       false,
	}
	prevSpecificPricing := models.SpecificPricing{}
	if err := sm.db.Model(&prevSpecificPricing).Preload(clause.Associations).Where("main_pricing_id = ?", mainPricingId).Last(&prevSpecificPricing).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "previous specific pricing"})
		return
	}
	if err := sm.db.Create(&newSpecificPricing).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing"})
		return
	}

	newProductExtraBuyerTypes := []models.ProductExtraBuyerTypes{}
	for _, prevProductExtraBuyerType := range prevSpecificPricing.ProductExtraBuyerTypes {
		newProductExtraBuyerType := models.ProductExtraBuyerTypes{
			ProductId:   2,
			ExtraId:     prevProductExtraBuyerType.ExtraId,
			PricingId:   newSpecificPricing.PricingId,
			BuyerTypeId: prevProductExtraBuyerType.BuyerTypeId,
			Price:       prevProductExtraBuyerType.Price,
			HasDiscount: prevProductExtraBuyerType.HasDiscount,
		}
		newProductExtraBuyerTypes = append(newProductExtraBuyerTypes, newProductExtraBuyerType)
	}
	err = sm.db.Create(&newProductExtraBuyerTypes).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing product extra buyer type"})
		return
	}
	newProductVenueBuyerTypes := []models.ProductVenueBuyerTypes{}
	for _, prevVenueExtraBuyerType := range prevSpecificPricing.ProductVenueBuyerTypes {
		newProductVenueBuyerType := models.ProductVenueBuyerTypes{
			ProductId:   2,
			VenueId:     prevVenueExtraBuyerType.VenueId,
			PricingId:   newSpecificPricing.PricingId,
			BuyerTypeId: prevVenueExtraBuyerType.BuyerTypeId,
			Price:       prevVenueExtraBuyerType.Price,
			HasDiscount: prevVenueExtraBuyerType.HasDiscount,
		}
		newProductVenueBuyerTypes = append(newProductVenueBuyerTypes, newProductVenueBuyerType)
	}

	if err := sm.db.Create(&newProductVenueBuyerTypes).Error; err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing product venue buyer type"})
		return
	}
	if req.RecurrentTime != nil {
		recurrentTimes := models.RecurrentTime{
			PricingId: newSpecificPricing.PricingId,
			Minutes:   req.RecurrentTime.Minutes,
			Hours:     req.RecurrentTime.Hours,
		}

		if err := sm.db.Create(&recurrentTimes).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing recurrent times"})
			return
		}
	}
	if req.DynamicPricingConfiguration != nil {
		dynamicPricingConfigs := models.DynamicPricingConfiguration{
			PricingId: newSpecificPricing.PricingId,
			Type:      req.DynamicPricingConfiguration.Type,
			StartHour: req.DynamicPricingConfiguration.StartHour,
			EndHour:   req.DynamicPricingConfiguration.EndHour,
		}

		if err := sm.db.Create(&dynamicPricingConfigs).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing dynamic pricings"})
			return
		}
		for _, occupancyRangeReq := range req.DynamicPricingConfiguration.OccupancyRanges {
			occupancyRange := models.OccupancyRange{DynamicPricingConfigurationId: dynamicPricingConfigs.DynamicPricingConfigurationId, Start: occupancyRangeReq.Start, End: occupancyRangeReq.End}
			if err := sm.db.Create(&occupancyRange).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing dynamic pricings occupancy range"})
				return
			}
		}
	}
}

func (sm *ServiceManager) PutPricingsConfiguration(w http.ResponseWriter, r *http.Request) {
	var reqSpecific = dto.SpecificPricingDTO{}
	var reqBase = dto.BasePricingPost{}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to read request body"})
		return
	}

	// Attempt decoding as both types
	_ = json.Unmarshal(bodyBytes, &reqSpecific)
	_ = json.Unmarshal(bodyBytes, &reqBase)

	// Heuristic: determine which one it actually is
	hasSpecificFields := len(reqSpecific.WeekDays) > 0 ||
		reqSpecific.StartHour != nil ||
		reqSpecific.EndHour != nil ||
		reqSpecific.RecurrentTime != nil ||
		reqSpecific.DynamicPricingConfiguration != nil

	hasBaseFields := reqBase.Color != "" && reqBase.StartDate != "" && reqBase.EndDate != ""

	if !hasBaseFields && !hasSpecificFields {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	if hasSpecificFields {
		fmt.Println("Updated specific pricing")
		specificPricing := models.SpecificPricing{}
		if err := helpers.GetById(&specificPricing, r, sm.db); err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Specific pricing not found"})
			return
		}

		newSpecificPricingUpdate := models.SpecificPricing{
			Name:         reqSpecific.Name,
			Weekdays:     reqSpecific.WeekDays,
			EnabledDates: reqSpecific.EnabledDates,
			StartHour:    reqSpecific.StartHour,
			EndHour:      reqSpecific.EndHour,
		}
		if err := sm.db.Model(&specificPricing).Select("Name", "Weekdays", "EnabledDates", "StartHour", "EndHour").Updates(newSpecificPricingUpdate).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing"})
			return
		}

		if reqSpecific.RecurrentTime != nil {
			prevRecurrentTime := models.RecurrentTime{}
			recurrentTimesUpdate := models.RecurrentTime{
				Minutes: reqSpecific.RecurrentTime.Minutes,
				Hours:   reqSpecific.RecurrentTime.Hours,
			}
			if err := sm.db.First(&prevRecurrentTime, "pricing_id = ?", specificPricing.PricingId).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Recurrent time not properly updated" + err.Error()})
				return
			}
			if err := sm.db.Model(&prevRecurrentTime).Updates(recurrentTimesUpdate).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Recurrent time not properly updated" + err.Error()})
				return
			}

		}
		if reqSpecific.DynamicPricingConfiguration != nil {
			dynamicPricingConfigsUpdate := models.DynamicPricingConfiguration{
				PricingId: specificPricing.PricingId,
				Type:      reqSpecific.DynamicPricingConfiguration.Type,
				StartHour: reqSpecific.DynamicPricingConfiguration.StartHour,
				EndHour:   reqSpecific.DynamicPricingConfiguration.EndHour,
			}
			prevDynamicPricingConfig := models.DynamicPricingConfiguration{}
			if err := sm.db.First(&prevDynamicPricingConfig, "pricing_id = ?", specificPricing.PricingId).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "DynamicpricingConfig not properly updated" + err.Error()})
				return
			}
			if err := sm.db.Model(&prevDynamicPricingConfig).Updates(dynamicPricingConfigsUpdate).Error; err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "DynamicpricingConfig not properly updated" + err.Error()})
				return
			}
			sm.db.Model(&prevDynamicPricingConfig).Association("OccupancyRanges").Clear()
			for _, occupancyRangereqSpecific := range reqSpecific.DynamicPricingConfiguration.OccupancyRanges {
				occupancyRange := models.OccupancyRange{DynamicPricingConfigurationId: prevDynamicPricingConfig.DynamicPricingConfigurationId, Start: occupancyRangereqSpecific.Start, End: occupancyRangereqSpecific.End}
				if err := sm.db.Create(&occupancyRange).Error; err != nil {
					helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "new specific pricing dynamic pricings occupancy range"})
					return
				}
			}
		}
	} else if hasBaseFields {
		fmt.Println("Updated main pricing")
		mainPricing := models.MainPricing{}
		if err := helpers.GetById(&mainPricing, r, sm.db); err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Main pricing not found"})
			return
		}

		mainPricingUpdate := models.MainPricing{
			Name:      reqBase.Name,
			Color:     reqBase.Color,
			StartDate: reqBase.StartDate,
			EndDate:   reqBase.EndDate,
		}
		if err := sm.db.Model(&mainPricing).Select("Name", "Color", "StartDate", "StartHour", "EndDate").Updates(mainPricingUpdate).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error() + " " + "cant update base pricing"})
			return
		}
	}

}

func (sm *ServiceManager) PutPricingsPriorities(w http.ResponseWriter, r *http.Request) {
	var req = dto.UpdatePricingPriorityDTO{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, pricing := range req {
		specificPricing := models.SpecificPricing{}
		if err := sm.db.First(&specificPricing, pricing.PricingId).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Specific pricing not found"})
			return
		}
		if err := sm.db.Model(&specificPricing).Update("Priority", pricing.Priority).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Specific pricing priority not updated"})
			return
		}
	}
}
func (sm *ServiceManager) PutPricingsTariffs(w http.ResponseWriter, r *http.Request) {
	var req = dto.UpdatePricingValuesDTO{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, pricing := range req {
		specificPricing := models.SpecificPricing{}

		if err := sm.db.First(&specificPricing, pricing.PricingId).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Specific pricing not found"})
			return
		}
		for _, venueReq := range pricing.ProductVenueBuyerTypes {
			venue := models.ProductVenueBuyerTypes{}
			venueUpdate := models.ProductVenueBuyerTypes{
				Price: venueReq.Price,
			}
			err := sm.db.Model(&venue).Find(&venue, venueReq.ProductVenueBuyerTypeId).Error
			if err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "venue pricing not found"})
				return
			}
			err = sm.db.Model(&venue).Updates(venueUpdate).Error
			if err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "venue pricing not updated"})
				return
			}
		}
		for _, extraReq := range pricing.ProductExtraBuyerTypes {
			extra := models.ProductExtraBuyerTypes{}
			extraUpdate := models.ProductExtraBuyerTypes{
				Price: extraReq.Price,
			}
			err := sm.db.Model(&extra).Find(&extra, extraReq.ProductExtraBuyerTypeId).Error
			if err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "extra pricing not found"})
				return
			}
			err = sm.db.Model(&extra).Updates(extraUpdate).Error
			if err != nil {
				helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "extra pricing not updated"})
				return
			}
		}
	}
}

// func (sm *ServiceManager) PutPricings(w http.ResponseWriter, r *http.Request) {
// 	var req []dto.PricingDTO
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
// 		return
// 	}
// 	err := createUpdateMainPricing(req, sm)
// 	if err != nil {
// 		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
// 		return
// 	}
// }

// func createUpdateMainPricing(req []dto.PricingDTO, sm *ServiceManager) error {
// 	for index, pricingReq := range req {
// 		pricing := models.MainPricing{}
// 		_, err := time.Parse("2006-01-02T00:00:00Z", pricingReq.StartDatetime)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = time.Parse("2006-01-02T15:04:05Z", pricingReq.StartDatetime)
// 		if err != nil {
// 			return err
// 		}

// 		if errors.Is(err, gorm.ErrRecordNotFound) || pricingReq.MainPricingId == nil {
// 			if index > 0 {
// 				prevMainPricing := models.MainPricing{}
// 				prevPricingReq := req[index-1]
// 				prevPricingId := *prevPricingReq.MainPricingId
// 				err = sm.db.Model(&prevMainPricing).Preload(clause.Associations).Where("main_pricing_id = ?", prevPricingId).Find(&prevMainPricing).Error
// 				if err != nil {
// 					return err
// 				}
// 				pricingUpdate := models.MainPricing{
// 					StartDate: pricingReq.StartDatetime,
// 					EndDate:   pricingReq.EndDatetime,
// 					ProductId: prevMainPricing.ProductId,
// 				}
// 				err = sm.db.Create(&pricingUpdate).Error
// 				if err != nil {
// 					return err
// 				}
// 				err = createBaseSpecificPricing(pricingUpdate, prevMainPricing, sm)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		} else if pricingReq.MainPricingId != nil {
// 			// err = sm.db.Model(&pricing).Where("start_datetime = ?", startDateTimeDb).Where("end_date_time = ?", endDateTimeDb).Find(&pricing).Error
// 			pricingUpdate := models.MainPricing{
// 				StartDate: pricingReq.StartDatetime,
// 				EndDate:   pricingReq.EndDatetime,
// 			}
// 			mainPricingId := *pricingReq.MainPricingId
// 			err = sm.db.Model(&pricing).Find(&pricing, mainPricingId).Error
// 			if err != nil {
// 				return err
// 			}
// 			err = sm.db.Model(&pricing).Updates(pricingUpdate).Error
// 			if err == nil {
// 				err := createUpdateSpecificPricing(pricing, pricingReq, sm)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func createBaseSpecificPricing(pricing models.MainPricing, prevMainPricing models.MainPricing, sm *ServiceManager) error {
	basePricing := models.SpecificPricing{
		MainPricingId: pricing.MainPricingId,
		Name:          "Tarifario Base",
		Priority:      1,
		Weekdays:      make([]int, 0),
		EnabledDates:  make([]string, 0),
		Default:       true,
	}
	err := sm.db.Create(&basePricing).Error
	if err != nil {
		return err
	}
	prevBasePricing := models.SpecificPricing{}
	for index := range prevMainPricing.Pricings {
		specificPricing := prevMainPricing.Pricings[index]
		if specificPricing.Default {
			prevBasePricing = prevMainPricing.Pricings[index]
			break
		}
	}
	err = sm.db.Model(&prevBasePricing).Preload(clause.Associations).Where("pricing_id = ?", prevBasePricing.PricingId).Find(&prevBasePricing).Error
	if err != nil {
		return err
	}
	newProductExtraBuyerTypes := []models.ProductExtraBuyerTypes{}
	for _, prevProductExtraBuyerType := range prevBasePricing.ProductExtraBuyerTypes {
		newProductExtraBuyerType := models.ProductExtraBuyerTypes{
			ProductId:   prevProductExtraBuyerType.ProductId,
			ExtraId:     prevProductExtraBuyerType.ExtraId,
			PricingId:   basePricing.PricingId,
			BuyerTypeId: prevProductExtraBuyerType.BuyerTypeId,
			Price:       prevProductExtraBuyerType.Price,
			HasDiscount: prevProductExtraBuyerType.HasDiscount,
		}
		newProductExtraBuyerTypes = append(newProductExtraBuyerTypes, newProductExtraBuyerType)
	}
	err = sm.db.Create(&newProductExtraBuyerTypes).Error
	if err != nil {
		return err
	}
	newProductVenueBuyerTypes := []models.ProductVenueBuyerTypes{}
	for _, prevVenueExtraBuyerType := range prevBasePricing.ProductVenueBuyerTypes {
		newProductVenueBuyerType := models.ProductVenueBuyerTypes{
			ProductId:   prevVenueExtraBuyerType.ProductId,
			VenueId:     prevVenueExtraBuyerType.VenueId,
			PricingId:   basePricing.PricingId,
			BuyerTypeId: prevVenueExtraBuyerType.BuyerTypeId,
			Price:       prevVenueExtraBuyerType.Price,
			HasDiscount: prevVenueExtraBuyerType.HasDiscount,
		}
		newProductVenueBuyerTypes = append(newProductVenueBuyerTypes, newProductVenueBuyerType)
	}
	err = sm.db.Create(&newProductVenueBuyerTypes).Error
	if err != nil {
		return err
	}
	return nil
}

// func createUpdateSpecificPricing(pricing models.MainPricing, reqItem dto.PricingDTO, sm *ServiceManager) error {
// 	for index, specificPricingReq := range reqItem.Pricings {
// 		if specificPricingReq.PricingId != nil {
// 			specificPricing := models.SpecificPricing{}
// 			specificPricingUpdate := models.SpecificPricing{
// 				Name:         specificPricingReq.Name,
// 				Priority:     specificPricingReq.Priority,
// 				Weekdays:     specificPricingReq.WeekDays,
// 				EnabledDates: specificPricingReq.EnabledDates,
// 			}
// 			pricingId := *specificPricingReq.PricingId
// 			err := sm.db.Model(&specificPricing).Preload(clause.Associations).Find(&specificPricing, pricingId).Error
// 			if err != nil {
// 				return err
// 			}
// 			err = sm.db.Model(&specificPricing).Updates(specificPricingUpdate).Error
// 			if err != nil {
// 				return err
// 			}
// 			for _, venueReq := range specificPricingReq.ProductVenueBuyerTypes {
// 				venue := models.ProductVenueBuyerTypes{}
// 				venueUpdate := models.ProductVenueBuyerTypes{
// 					Price: venueReq.Price,
// 				}
// 				err := sm.db.Model(&venue).Find(&venue, venueReq.ProductVenueBuyerTypeId).Error
// 				if err != nil {
// 					return err
// 				}
// 				err = sm.db.Model(&venue).Updates(venueUpdate).Error
// 				if err != nil {
// 					return err
// 				}
// 			}
// 			for _, extraReq := range specificPricingReq.ProductExtraBuyerTypes {
// 				extra := models.ProductExtraBuyerTypes{}
// 				extraUpdate := models.ProductExtraBuyerTypes{
// 					Price: extraReq.Price,
// 				}
// 				err := sm.db.Model(&extra).Find(&extra, extraReq.ProductExtraBuyerTypeId).Error
// 				if err != nil {
// 					return err
// 				}
// 				err = sm.db.Model(&extra).Updates(extraUpdate).Error
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		} else {
// 			fmt.Println(specificPricingReq)
// 			newSpecificPricing := models.SpecificPricing{
// 				Name:          specificPricingReq.Name,
// 				Priority:      specificPricingReq.Priority,
// 				Weekdays:      specificPricingReq.WeekDays,
// 				EnabledDates:  specificPricingReq.EnabledDates,
// 				MainPricingId: pricing.MainPricingId,
// 			}
// 			err := sm.db.Create(&newSpecificPricing).Error
// 			if err != nil {
// 				return err
// 			}
// 			if index > 0 {
// 				prevPricingReq := reqItem.Pricings[0]
// 				prevPricing := models.SpecificPricing{}
// 				err := sm.db.Model(&prevPricing).Preload(clause.Associations).Where("pricing_id = ?", *prevPricingReq.PricingId).Find(&prevPricing).Error
// 				if err != nil {
// 					return err
// 				}
// 				newProductExtraBuyerTypes := []models.ProductExtraBuyerTypes{}
// 				fmt.Println(pricing.ProductId)
// 				for _, prevProductExtraBuyerType := range prevPricing.ProductExtraBuyerTypes {
// 					newProductExtraBuyerType := models.ProductExtraBuyerTypes{
// 						ProductId:   pricing.ProductId,
// 						ExtraId:     prevProductExtraBuyerType.ExtraId,
// 						PricingId:   newSpecificPricing.PricingId,
// 						BuyerTypeId: prevProductExtraBuyerType.BuyerTypeId,
// 						Price:       prevProductExtraBuyerType.Price,
// 						HasDiscount: prevProductExtraBuyerType.HasDiscount,
// 					}
// 					newProductExtraBuyerTypes = append(newProductExtraBuyerTypes, newProductExtraBuyerType)
// 				}
// 				err = sm.db.Create(&newProductExtraBuyerTypes).Error
// 				if err != nil {
// 					return err
// 				}
// 				newProductVenueBuyerTypes := []models.ProductVenueBuyerTypes{}
// 				for _, prevVenueExtraBuyerType := range prevPricing.ProductVenueBuyerTypes {
// 					newProductVenueBuyerType := models.ProductVenueBuyerTypes{
// 						ProductId:   pricing.ProductId,
// 						VenueId:     prevVenueExtraBuyerType.VenueId,
// 						PricingId:   newSpecificPricing.PricingId,
// 						BuyerTypeId: prevVenueExtraBuyerType.BuyerTypeId,
// 						Price:       prevVenueExtraBuyerType.Price,
// 						HasDiscount: prevVenueExtraBuyerType.HasDiscount,
// 					}
// 					newProductVenueBuyerTypes = append(newProductVenueBuyerTypes, newProductVenueBuyerType)
// 				}
// 				err = sm.db.Create(&newProductVenueBuyerTypes).Error
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }
