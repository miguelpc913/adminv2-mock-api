package services

import (
	"encoding/json"
	"net/http"

	dtoBoxOffice "github.com/tiqueteo/adminv2-mock-api/api/dto/boxOffice"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (sm *ServiceManager) GetBO(w http.ResponseWriter, r *http.Request) {
	var boxOffices []models.BoxOffice
	response := helpers.PaginateRequest(r, boxOffices, sm.db, "boxOffices", "status")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetBOById(w http.ResponseWriter, r *http.Request) {
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no entity with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, boxOffice)
	}
}

func (sm *ServiceManager) PutBOBasicConfigurations(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficeBasicConfigurations{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOCashCount(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficeCashCount{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOPresentations(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficePresentations{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOFunctionalities(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficeFunctionalities{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOLanguages(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficeLanguages{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	var allowedAppLanguages []models.AllowedAppLanguages
	for _, language := range req.AllowedAppLanguages {
		newAppLanguage := models.AllowedAppLanguages{
			LanguageCode: language.LanguageCode,
			DisplayOrder: language.DisplayOrder,
			BoxOfficeId:  boxOffice.BoxOfficeId,
		}
		allowedAppLanguages = append(allowedAppLanguages, newAppLanguage)
	}
	err = sm.db.Create(&allowedAppLanguages).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	sm.db.Model(&boxOffice).Association("AllowedAppLanguages").Replace(allowedAppLanguages)

	var allowedTicketLanguages []models.AllowedTicketLanguages
	for _, language := range req.AllowedTicketLanguages {
		newAppLanguage := models.AllowedTicketLanguages{
			LanguageCode: language.LanguageCode,
			DisplayOrder: language.DisplayOrder,
			BoxOfficeId:  boxOffice.BoxOfficeId,
		}
		allowedTicketLanguages = append(allowedTicketLanguages, newAppLanguage)
	}
	err = sm.db.Create(&allowedTicketLanguages).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	sm.db.Model(&boxOffice).Association("AllowedTicketLanguages").Replace(allowedTicketLanguages)
	boxOfficeUpdate := models.BoxOffice{
		LanguageCode: req.LanguageCode,
	}
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOPrintSettings(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficePrintSettings{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := models.BoxOffice{
		PrintTicket:                  req.PrintTicket,
		OptionalPrintTicket:          req.OptionalPrintTicket,
		PrintTicketPrice:             req.PrintTicketPrice,
		OptionalPrintTicketPrice:     req.OptionalPrintTicketPrice,
		PrintSummary:                 req.PrintSummary,
		OptionalPrintSummary:         req.OptionalPrintSummary,
		HighlightPrintedReservations: req.HighlightPrintedReservations,
		SingleDocPrint:               req.SingleDocPrint,
		AllowedTicketGroupTypes:      req.AllowedTicketGroupTypes,
		PrintCashCount:               req.PrintCashCount,
	}
	err = sm.db.Model(&boxOffice).Select("PrintTicket", "OptionalPrintTicket", "PrintTicketPrice", "OptionalPrintTicketPrice", "PrintSummary", "OptionalPrintSummary", "HighlightPrintedReservations", "SingleDocPrint", "AllowedTicketGroupTypes", "PrintCashCount").Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOPaymentSettings(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficePaymentRequest{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOValidations(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficeValidations{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOAdvancedSettings(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficeAdvancedSettings{}
	boxOffice := models.BoxOffice{}
	err := helpers.GetById(&boxOffice, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no box office with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	boxOfficeUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&boxOffice).Updates(boxOfficeUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}

func (sm *ServiceManager) PutBOSalesGroups(w http.ResponseWriter, r *http.Request) {
	var boxOffice models.BoxOffice
	err := helpers.UpdateRelation(r, boxOffice, models.SalesGroup{}, sm.db, "SalesGroupSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PutBOProducts(w http.ResponseWriter, r *http.Request) {
	var boxOffice models.BoxOffice
	err := helpers.UpdateRelation(r, boxOffice, models.Product{}, sm.db, "ProductSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PostBO(w http.ResponseWriter, r *http.Request) {
	req := &dtoBoxOffice.BoxOfficePost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	boxOffice := models.BoxOffice{}
	err := helpers.CopyStructFields(req, &boxOffice)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = sm.db.Create(&boxOffice).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, boxOffice)
}
