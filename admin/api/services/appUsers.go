package services

import (
	"encoding/json"
	"errors"
	"net/http"

	appUserDto "github.com/tiqueteo/adminv2-mock-api/api/dto/appUser"
	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
	"gorm.io/gorm"
)

func (sm *ServiceManager) GetAppUsers(w http.ResponseWriter, r *http.Request) {
	var appUsers []models.User
	response := helpers.PaginateRequest(r, appUsers, sm.db, "appUsers")
	helpers.WriteJSON(w, http.StatusOK, response)
}

func (sm *ServiceManager) GetAppUserById(w http.ResponseWriter, r *http.Request) {
	appUser := models.User{}
	err := helpers.GetById(&appUser, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no appUser with that id"})
	} else {
		helpers.WriteJSON(w, http.StatusOK, appUser)
	}
}

func (sm *ServiceManager) PutAppUserIdentity(w http.ResponseWriter, r *http.Request) {
	req := &appUserDto.AppUserIdentityDTO{}
	appUser := models.User{}
	err := helpers.GetById(&appUser, r, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "There is no appUser with that id"})
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	appUserUpdate := helpers.StructToMap(req)
	err = sm.db.Model(&appUser).Updates(appUserUpdate).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Couldn't update"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, appUserUpdate)
}

func (sm *ServiceManager) PutReportSet(w http.ResponseWriter, r *http.Request) {
	appUser := models.User{}
	err := helpers.UpdateRelation(r, appUser, models.Report{}, sm.db, "ReportSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}
func (sm *ServiceManager) PutPointOfSaleSet(w http.ResponseWriter, r *http.Request) {
	appUser := models.User{}
	err := helpers.UpdateRelation(r, appUser, models.PointsOfSale{}, sm.db, "PointOfSaleSet")
	if err != nil {
		helpers.WriteJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	} else {
		helpers.WriteJSON(w, http.StatusOK, map[string]string{"Success": "Updated properly"})
	}
}

func (sm *ServiceManager) PostValidateUser(w http.ResponseWriter, r *http.Request) {
	type userNameValidation struct {
		username string
	}
	req := &userNameValidation{}
	appUser := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	if err := sm.db.First(&appUser, "user_name = ?", req.username).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		helpers.WriteJSON(w, http.StatusOK, map[string]bool{"valid": true})
	} else {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]bool{"valid": false})
	}
}

func (sm *ServiceManager) PostAppUser(w http.ResponseWriter, r *http.Request) {

	req := &appUserDto.AppUserCreateDTO{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
	pointsOfSale := []models.PointsOfSale{}
	err := helpers.GetByIds(&pointsOfSale, req.PointOfSaleSet, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "point of sales are not valid"})
		return
	}
	reports := []models.Report{}
	err = helpers.GetByIds(&reports, req.ReportSet, sm.db)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "reports are not valid"})
		return
	}

	affiliateAgreement := models.User{
		Status:         req.Status,
		Type:           models.UserType(req.Type),
		Profile:        models.UserProfile(req.Profile),
		Name:           req.Name,
		LastName:       req.LastName,
		UserName:       req.UserName,
		Email:          req.Email,
		Password:       req.Password,
		ReportSet:      reports,
		PointOfSaleSet: pointsOfSale,
	}
	err = sm.db.Create(&affiliateAgreement).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, affiliateAgreement)
}
