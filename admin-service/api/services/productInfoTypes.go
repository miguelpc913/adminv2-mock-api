package services

import (
	"net/http"

	"github.com/tiqueteo/adminv2-mock-api/api/helpers"
	"github.com/tiqueteo/adminv2-mock-api/db/models"
)

func (serviceManager *ServiceManager) GetProductInfoType(w http.ResponseWriter, r *http.Request) {
	var productInfoTypes []models.ProductInfoType
	err := serviceManager.db.Find(&productInfoTypes).Error
	if err != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error in db query"})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, productInfoTypes)
}
