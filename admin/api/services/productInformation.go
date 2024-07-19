package services

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/go-chi/chi"
	dtoPi "github.com/tiqueteo/adminv2-mock-api/api/dto/productInformation"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm/clause"
)

func (serviceManager *ServiceManager) GetProductInfos(w http.ResponseWriter, r *http.Request) {
	var productsInformation []models.ProductInfo
	pagination := helpers.ManagePaginationQueries(r)
	response := make(map[string]interface{})
	err := serviceManager.db.Preload("ProductInfoType").Model(&productsInformation).Count(&pagination.TotalItems).Limit(pagination.Limit).Offset(pagination.Offset).Find(&productsInformation).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error in pagination"})
		return
	}
	for index := range productsInformation {
		infoTypeDescription := productsInformation[index].ProductInfoType.Id
		productsInformation[index].InfoType = infoTypeDescription
	}
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(pagination.TotalItems) / float64(pagination.Limit)))
	response["totalItems"] = pagination.TotalItems
	response["productInfos"] = productsInformation
	helpers.WriteJSON(w, http.StatusOK, response)
}
func (sm *ServiceManager) GetProductInfoById(w http.ResponseWriter, r *http.Request) {
	var productInfo models.ProductInfo
	id := chi.URLParam(r, "id")
	err := sm.db.Preload(clause.Associations).Find(&productInfo, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not product info with that id"})
		return
	}
	SalesGroupSet := []int{}
	for _, salesGroup := range productInfo.SalesGroupSet {
		SalesGroupSet = append(SalesGroupSet, salesGroup.SalesGroupId)
	}
	ProductSet := []int{}
	for _, product := range productInfo.ProductSet {
		ProductSet = append(ProductSet, product.ProductId)
	}
	VenueSet := []int{}
	for _, venue := range productInfo.VenueSet {
		VenueSet = append(VenueSet, venue.VenueCapacityId)
	}
	productInfoResp := dtoPi.GetProductInfo{
		VenueSet:      VenueSet,
		ProductSet:    ProductSet,
		SalesGroupSet: SalesGroupSet,
		InfoType:      productInfo.ProductInfoType.Id,
		Status:        productInfo.Status,
		Name:          productInfo.Name,
		InternalName:  productInfo.InternalName,
		Description:   productInfo.Description,
		Icon:          productInfo.Icon,
		CalendarColor: productInfo.CalendarColor,
		Weekdays:      productInfo.Weekdays,
		SelectedDates: productInfo.SelectedDates,
		DisplayOrder:  productInfo.DisplayOrder,
	}
	helpers.WriteJSON(w, http.StatusOK, productInfoResp)
}

func (sm *ServiceManager) PutProductInfoIdentity(w http.ResponseWriter, r *http.Request) {
	var productInformation models.ProductInfo
	var req dtoPi.PutProductInfoIdentity
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	err := sm.db.Model(productInformation).Where("product_info_id = ?", id).Updates(req).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, productInformation)
}

func (sm *ServiceManager) PutProductInfoSettings(w http.ResponseWriter, r *http.Request) {
	var productInformation models.ProductInfo
	var req dtoPi.PutProductInfoSettings
	id := chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for index := range req.SelectedDates {
		_, err := helpers.ParseDate(req.SelectedDates[index])
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Selected dates is not valid"})
			return
		}
	}
	infoType := models.ProductInfoType{}
	err := sm.db.First(&infoType, "id = ?", req.InfoType).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Info type is not valid"})
		return
	}
	productInfo := models.ProductInfo{
		Name:              req.Name,
		Description:       req.Description,
		Icon:              req.Icon,
		InfoType:          req.InfoType,
		CalendarColor:     req.CalendarColor,
		Weekdays:          req.Weekdays,
		SelectedDates:     req.SelectedDates,
		ProductInfoType:   infoType,
		ProductInfoTypeId: infoType.ProductInfoTypeId,
	}
	err = sm.db.Model(&productInformation).Where("product_info_id = ?", id).Updates(productInfo).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, productInformation)
}

func (sm *ServiceManager) PutProductInfoSalesGroups(w http.ResponseWriter, r *http.Request) {
	var productInfo models.ProductInfo
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&productInfo, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not product info with that id"})
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

	sm.db.Model(&productInfo).Association("SalesGroupSet").Replace(salesGroups)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutProductInfoProducts(w http.ResponseWriter, r *http.Request) {
	var productInfo models.ProductInfo
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&productInfo, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not product info with that id"})
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

	sm.db.Model(&productInfo).Association("ProductSet").Replace(products)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutProductInfoVenue(w http.ResponseWriter, r *http.Request) {
	var productInfo models.ProductInfo
	req := []int{}
	id := chi.URLParam(r, "id")
	err := sm.db.Find(&productInfo, id).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is not product info with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	venues := []models.VenueCapacity{}
	for _, id := range req {
		venue := models.VenueCapacity{}
		if err := sm.db.First(&venue, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Products are not valid"})
			return
		}
		venues = append(venues, venue)
	}

	sm.db.Model(&productInfo).Association("VenueSet").Replace(venues)

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
}

func (sm *ServiceManager) PutOrderProductInfos(w http.ResponseWriter, r *http.Request) {
	var productsInformation []models.ProductInfo
	var req []dtoPi.DisplayOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	for _, item := range req {
		err := sm.db.Model(productsInformation).Where("product_info_id = ?", item.ProductInfoId).Update("display_order", item.DisplayOrder).Error
		if err != nil {
			log.Printf("failed to update product info id %d: %v", item.ProductInfoId, err)
		}
	}

	helpers.WriteJSON(w, http.StatusOK, productsInformation)
}

func (sm *ServiceManager) PostProductInfos(w http.ResponseWriter, r *http.Request) {
	infoReq := &dtoPi.PostProductInfo{}
	if err := json.NewDecoder(r.Body).Decode(&infoReq); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	//Validate dates
	for index := range infoReq.SelectedDates {
		_, err := helpers.ParseDate(infoReq.SelectedDates[index])
		if err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Selected dates is not valid"})
			return
		}
	}

	//Manage associations
	salesGroups := []models.SalesGroup{}
	for _, id := range infoReq.SalesGroupSet {
		salesGroup := models.SalesGroup{}
		if err := sm.db.First(&salesGroup, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "SalesGroups are not valid"})
			return
		}
		salesGroups = append(salesGroups, salesGroup)
	}

	products := []models.Product{}
	for _, id := range infoReq.ProductSet {
		product := models.Product{}
		if err := sm.db.First(&product, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Products are not valid"})
			return
		}
		products = append(products, product)
	}

	venues := []models.VenueCapacity{}
	for _, id := range infoReq.VenueSet {
		venue := models.VenueCapacity{}
		if err := sm.db.First(&venue, id).Error; err != nil {
			helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Venues are not valid"})
			return
		}
		venues = append(venues, venue)
	}

	//Find info type
	infoType := models.ProductInfoType{}
	err := sm.db.First(&infoType, "id = ?", infoReq.InfoType).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Info type is not valid"})
		return
	}

	productInfo := models.ProductInfo{
		SalesGroupSet:     salesGroups,
		VenueSet:          venues,
		ProductSet:        products,
		Status:            infoReq.Status,
		Name:              infoReq.Name,
		InternalName:      infoReq.InternalName,
		Description:       infoReq.Description,
		Icon:              infoReq.Icon,
		InfoType:          infoReq.InfoType,
		CalendarColor:     infoReq.CalendarColor,
		Weekdays:          infoReq.Weekdays,
		SelectedDates:     infoReq.SelectedDates,
		ProductInfoType:   infoType,
		ProductInfoTypeId: infoType.ProductInfoTypeId,
	}
	err = sm.db.Create(&productInfo).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, productInfo)
}
