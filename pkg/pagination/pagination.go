package pagination

import (
	"math"
	"strconv"
	"strings"
	"users-api-server/pkg/consts"
)

func CreatePaginationRequestFromParams(query map[string][]string) Request {
	paginationRequest := Request{
		Page:  consts.DefaultPageNumber,
		Limit: consts.DefaultLimit,
		Sort:  consts.DefaultSort,
	}
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case consts.Limit:
			limit, err := strconv.Atoi(queryValue)
			if err == nil {
				paginationRequest.Limit = limit
			}
			break
		case consts.Page:
			page, err := strconv.Atoi(queryValue)
			if err == nil {
				paginationRequest.Page = page
			}
			break
		case consts.Sort:
			if queryValue != "" {
				paginationRequest.Sort = queryValue
			}
			break
		case consts.Filter:
			if len(value) != 0 {
				paginationRequest.Filter = make(map[string]interface{}, len(value))
				for _, v := range value {
					splits := strings.Split(v, "=")
					paginationRequest.Filter[splits[0]] = splits[1]
				}
			}
			break
		}
	}
	return paginationRequest
}

func CountTotalPagesAndNextPage(totalRecords int64, itemsPerPage, pageNumber int) (int64, int64) {
	totalPages := math.Round(float64(totalRecords) / float64(itemsPerPage))
	if totalPages == 0 {
		totalPages = 1
	}
	nextPage := pageNumber + 1
	if float64(nextPage) >= totalPages {
		nextPage = int(totalPages)
	}

	return int64(totalPages), int64(nextPage)
}

type Request struct {
	Page   int
	Limit  int
	Sort   string
	Filter map[string]interface{}
}
