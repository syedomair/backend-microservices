package request

import (
	"errors"
	"net/http"
	"strconv"
)

// QueryParams represents query parameters for pagination and sorting.
type QueryParams struct {
	Limit   int
	Page    int
	OrderBy string
	Sort    string
}

// ValidateQueryString validates query string parameters and returns QueryParams.
func ValidateQueryString(r *http.Request, defaultLimit string, defaultPage string, defaultOrderby string, defaultSort string) (QueryParams, error) {
	params := QueryParams{}

	// Extract query parameters
	limit, err := getLimit(r, defaultLimit)
	if err != nil {
		return params, err
	}
	page, err := getPage(r, defaultPage)
	if err != nil {
		return params, err
	}
	orderby, err := getOrderBy(r, defaultOrderby)
	if err != nil {
		return params, err
	}
	sort, err := getSort(r, defaultSort)
	if err != nil {
		return params, err
	}

	// Adjust page offset
	if page != 0 {
		page = (limit * page) - limit
	}
	return QueryParams{
		Limit:   limit,
		Page:    page,
		OrderBy: orderby,
		Sort:    sort,
	}, nil
}

// getLimit retrieves and validates the 'limit' query parameter.
func getLimit(r *http.Request, defaultLimit string) (int, error) {
	limits, ok := r.URL.Query()["limit"]
	if !ok || len(limits[0]) < 1 {
		return parseInt(defaultLimit)
	}
	return parseInt(limits[0])
}

// getPage retrieves and validates the 'page' query parameter.
func getPage(r *http.Request, defaultPage string) (int, error) {
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) < 1 {
		return parseInt(defaultPage)
	}
	return parseInt(pages[0])
}

// getOrderBy retrieves and validates the 'orderby' query parameter.
func getOrderBy(r *http.Request, defaultOrderby string) (string, error) {
	orderbys, ok := r.URL.Query()["orderby"]
	if !ok || len(orderbys[0]) < 1 {
		return defaultOrderby, nil
	}
	orderby := orderbys[0]
	if _, err := strconv.Atoi(orderby); err == nil {
		return "", errors.New("invalid 'orderby' value in query string. Must be a string. ")
	}
	return orderby, nil
}

// getSort retrieves and validates the 'sort' query parameter.
func getSort(r *http.Request, defaultSort string) (string, error) {
	sorts, ok := r.URL.Query()["sort"]
	if !ok || len(sorts[0]) < 1 {
		return defaultSort, nil
	}
	sort := sorts[0]
	if _, err := strconv.Atoi(sort); err == nil {
		return "", errors.New("invalid 'sort' value in query string. Must be a string. ")
	}
	if (sort != "asc") && (sort != "desc") {
		return "", errors.New("invalid 'sort' value in query string. Must be either 'asc' or 'desc'. ")
	}
	return sort, nil
}

// parseInt converts a string to an integer.
func parseInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("invalid number in query string. Must be a number. ")
	}
	return i, nil
}
