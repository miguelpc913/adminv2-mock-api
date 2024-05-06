package helpers

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit       int    `json:"limit"`
	CurrentPage int    `json:"currentPage"`
	Sort        string `json:"sort"`
}

func GeneratePaginationFromRequest(r *http.Request) Pagination {
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
