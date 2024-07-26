package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	dto "github.com/tiqueteo/adminv2-mock-api/api/dto/pricings"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (sm *ServiceManager) GetPricings(w http.ResponseWriter, r *http.Request) {
	pricings := []models.MainPricing{}
	productId := r.URL.Query().Get("productId")
	sm.db.Preload("Pricings.ProductExtraBuyerTypes").Preload("Pricings.ProductVenueBuyerTypes").Preload(clause.Associations).Model(&pricings).Where("product_id = ?", productId).Find(&pricings)
	helpers.WriteJSON(w, http.StatusOK, pricings)
}

func (sm *ServiceManager) PutPricings(w http.ResponseWriter, r *http.Request) {
	var req []dto.PricingDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := createUpdateMainPricing(req, sm)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
}

func createUpdateMainPricing(req []dto.PricingDTO, sm *ServiceManager) error {
	for index, pricingReq := range req {
		pricing := models.MainPricing{}
		_, err := time.Parse("2006-01-02T00:00:00Z", pricingReq.StartDatetime)
		if err != nil {
			return err
		}
		_, err = time.Parse("2006-01-02T15:04:05Z", pricingReq.StartDatetime)
		if err != nil {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) || pricingReq.MainPricingId == nil {
			if index > 0 {
				prevMainPricing := models.MainPricing{}
				fmt.Println(req[index-1])
				prevPricingReq := req[index-1]
				prevPricingId := *prevPricingReq.MainPricingId
				err = sm.db.Model(&prevMainPricing).Preload(clause.Associations).Where("main_pricing_id = ?", prevPricingId).Find(&prevMainPricing).Error
				if err != nil {
					return err
				}
				pricingUpdate := models.MainPricing{
					StartDatetime: pricingReq.StartDatetime,
					EndDateTime:   pricingReq.EndDatetime,
					ProductId:     prevMainPricing.ProductId,
				}
				err = sm.db.Create(&pricingUpdate).Error
				if err != nil {
					return err
				}
				err = createBaseSpecificPricing(pricingUpdate, prevMainPricing, sm)
				if err != nil {
					return err
				}
			}
		} else if pricingReq.MainPricingId != nil {
			// err = sm.db.Model(&pricing).Where("start_datetime = ?", startDateTimeDb).Where("end_date_time = ?", endDateTimeDb).Find(&pricing).Error
			pricingUpdate := models.MainPricing{
				StartDatetime: pricingReq.StartDatetime,
				EndDateTime:   pricingReq.EndDatetime,
			}
			mainPricingId := *pricingReq.MainPricingId
			err = sm.db.Model(&pricing).Find(&pricing, mainPricingId).Error
			if err != nil {
				return err
			}
			err = sm.db.Model(&pricing).Updates(pricingUpdate).Error
			if err == nil {
				err := createUpdateSpecificPricing(pricing, pricingReq, sm)
				if err != nil {
					return err
				}
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

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

func createUpdateSpecificPricing(pricing models.MainPricing, reqItem dto.PricingDTO, sm *ServiceManager) error {
	for index, specificPricingReq := range reqItem.Pricings {
		if specificPricingReq.PricingId != nil {
			specificPricing := models.SpecificPricing{}
			specificPricingUpdate := models.SpecificPricing{
				Name:         specificPricingReq.Name,
				Priority:     specificPricingReq.Priority,
				Weekdays:     specificPricingReq.WeekDays,
				EnabledDates: specificPricingReq.EnabledDates,
			}
			pricingId := *specificPricingReq.PricingId
			err := sm.db.Model(&specificPricing).Preload(clause.Associations).Find(&specificPricing, pricingId).Error
			if err != nil {
				return err
			}
			err = sm.db.Model(&specificPricing).Updates(specificPricingUpdate).Error
			if err != nil {
				return err
			}
			for _, venueReq := range specificPricingReq.ProductVenueBuyerTypes {
				venue := models.ProductVenueBuyerTypes{}
				venueUpdate := models.ProductVenueBuyerTypes{
					Price: venueReq.Price,
				}
				err := sm.db.Model(&venue).Find(&venue, venueReq.ProductVenueBuyerTypeId).Error
				if err != nil {
					return err
				}
				err = sm.db.Model(&venue).Updates(venueUpdate).Error
				if err != nil {
					return err
				}
			}
			for _, extraReq := range specificPricingReq.ProductExtraBuyerTypes {
				extra := models.ProductExtraBuyerTypes{}
				extraUpdate := models.ProductExtraBuyerTypes{
					Price: extraReq.Price,
				}
				err := sm.db.Model(&extra).Find(&extra, extraReq.ProductExtraBuyerTypeId).Error
				if err != nil {
					return err
				}
				err = sm.db.Model(&extra).Updates(extraUpdate).Error
				if err != nil {
					return err
				}
			}
		} else {
			fmt.Println(specificPricingReq)
			newSpecificPricing := models.SpecificPricing{
				Name:          specificPricingReq.Name,
				Priority:      specificPricingReq.Priority,
				Weekdays:      specificPricingReq.WeekDays,
				EnabledDates:  specificPricingReq.EnabledDates,
				MainPricingId: pricing.MainPricingId,
			}
			err := sm.db.Create(&newSpecificPricing).Error
			if err != nil {
				return err
			}
			if index > 0 {
				prevPricingReq := reqItem.Pricings[0]
				prevPricing := models.SpecificPricing{}
				err := sm.db.Model(&prevPricing).Preload(clause.Associations).Where("pricing_id = ?", *prevPricingReq.PricingId).Find(&prevPricing).Error
				if err != nil {
					return err
				}
				newProductExtraBuyerTypes := []models.ProductExtraBuyerTypes{}
				fmt.Println(pricing.ProductId)
				for _, prevProductExtraBuyerType := range prevPricing.ProductExtraBuyerTypes {
					newProductExtraBuyerType := models.ProductExtraBuyerTypes{
						ProductId:   pricing.ProductId,
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
					return err
				}
				newProductVenueBuyerTypes := []models.ProductVenueBuyerTypes{}
				for _, prevVenueExtraBuyerType := range prevPricing.ProductVenueBuyerTypes {
					newProductVenueBuyerType := models.ProductVenueBuyerTypes{
						ProductId:   pricing.ProductId,
						VenueId:     prevVenueExtraBuyerType.VenueId,
						PricingId:   newSpecificPricing.PricingId,
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
			}
		}
	}
	return nil
}
