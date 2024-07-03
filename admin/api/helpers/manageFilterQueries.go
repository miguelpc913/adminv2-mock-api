package helpers

import (
	"net/http"

	"gorm.io/gorm"
)

func manageFilterQueries(r *http.Request, query *gorm.DB) {
	urlQueries := r.URL.Query()
	for urlQueryKey, urlQueryValues := range urlQueries {
		if urlQueryKey != "size" && urlQueryKey != "page" && urlQueryKey != "sort" {
			for _, filterValue := range urlQueryValues {
				paramQuery := ToSnakeCase(urlQueryKey) + " = ?"
				if filterValue == "true" || filterValue == "false" {
					query = query.Where(paramQuery, filterValue == "true")
				} else {
					query = query.Where(paramQuery, filterValue)
				}
			}
		}
	}
}
