package helpers

import (
	"math"
	"net/http"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func generatePaginationFromRequest(r *http.Request) Pagination {
	size := 10
	page := 1
	sort := "created_at asc"
	query := r.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "size":
			size, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}
	return Pagination{
		Limit:       size,
		CurrentPage: page,
		Sort:        sort,
	}

}

func PaginateRequest[T any](r *http.Request, model T, db *gorm.DB, reponseKey string, filterParam string) map[string]interface{} {
	pagination := generatePaginationFromRequest(r)
	response := make(map[string]interface{})
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	var totalItems int64
	param := r.URL.Query().Get(filterParam)
	paramQuery := filterParam + " = ?"
	if param != "" {
		_ = db.Preload(clause.Associations).Model(&model).Where(paramQuery, param).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&model)
	} else {
		_ = db.Preload(clause.Associations).Model(&model).Count(&totalItems).Limit(pagination.Limit).Offset(offset).Find(&model)
	}
	response[reponseKey] = model
	response["limit"] = pagination.Limit
	response["currentPage"] = pagination.CurrentPage
	response["totalPages"] = int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
	response["totalItems"] = totalItems
	return response
}
