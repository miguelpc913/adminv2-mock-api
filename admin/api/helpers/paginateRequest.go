package helpers

import (
	"math"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PaginateRequest[T any](r *http.Request, model T, db *gorm.DB, reponseKey string) map[string]interface{} {
	pagination := ManagePaginationQueries(r)
	dbQuery := db.Preload(clause.Associations).Model(&model)
	manageFilterQueries(r, dbQuery)
	dbQuery.Count(&pagination.TotalItems).Limit(pagination.Limit).Offset(pagination.Offset).Find(&model)
	response := make(map[string]interface{})
	response[reponseKey] = model
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(pagination.TotalItems) / float64(pagination.Limit)))
	response["totalItems"] = pagination.TotalItems
	return response
}
