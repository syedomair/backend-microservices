package request

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	errorCodePrefix = "5"
)

const (
	SUCCESS = "success"
	FAILURE = "failure"
	//ERROR_UNEXPECTED = "Unexpected Error."
	STRING_SMALL    = "STRING_SMALL"
	STRING_LARGE    = "STRING_LARGE"
	STRING_EMAIL    = "STRING_EMAIL"
	STRING_PASSWORD = "STRING_PASSWORD"
	STRING_DATE     = "STRING_DATE"
	STRING_DATETIME = "STRING_DATETIME"
	INT             = "INT"
	INT_BOOLEAN     = "INT_BOOLEAN"
	ID              = "ID"
	EMAIL           = "EMAIL"
	PASSWORD        = "PASSWORD"
)

// GetRequestID Public
func GetRequestID(r *http.Request) string {
	requestID, _ := r.Context().Value("ContextKeyRequestID").(string)
	return requestID
}

// ValidateQueryString Public
func ValidateQueryString(r *http.Request, defaultLimit string, defaultPage string, defaultOrderby string, defaultSort string) (int, int, string, string, error) {

	strLimit := "0"
	page := ""
	orderby := ""
	sort := ""

	limits, ok := r.URL.Query()["limit"]
	if !ok || len(limits[0]) < 1 {
		strLimit = "0"
	} else {

		strLimit = limits[0]
	}

	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) < 1 {
		page = ""
	} else {
		page = pages[0]
	}
	orderbys, ok := r.URL.Query()["orderby"]
	if !ok || len(orderbys[0]) < 1 {
		orderby = ""
	} else {
		orderby = orderbys[0]
	}
	sorts, ok := r.URL.Query()["sort"]
	if !ok || len(sorts[0]) < 1 {
		sort = ""
	} else {
		sort = sorts[0]
	}

	if strLimit != "0" {
		if _, err := strconv.Atoi(strLimit); err != nil {
			return 0, 0, "", "", errors.New("invalid 'limit' number in query string. Must be a number. ")
		}
	} else {
		strLimit = defaultLimit
	}
	if page != "" {
		if _, err := strconv.Atoi(page); err != nil {
			return 0, 0, "", "", errors.New("invalid 'page' number in query string. Must be a number. ")
		}
	} else {
		page = defaultPage
	}

	intLimit, err := strconv.Atoi(strLimit)
	if err != nil {
		return 0, 0, "", "", errors.New("invalid 'limit' number in query string. Must be a number. ")
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, "", "", errors.New("invalid 'page' number in query string. Must be a number. ")
	}

	intPage = (intLimit * intPage) - intLimit

	if orderby != "" {
		if _, err := strconv.Atoi(orderby); err == nil {
			return 0, 0, "", "", errors.New("invalid 'orderby' value in query string. Must be a string. ")
		}
	} else {
		orderby = defaultOrderby
	}

	if sort != "" {
		if _, err := strconv.Atoi(sort); err == nil {
			return 0, 0, "", "", errors.New("invalid 'sort' value in query string. Must be a string. ")
		}
		if (sort != "asc") && (sort != "desc") {
			return 0, 0, "", "", errors.New("invalid 'sort' value in query string. Must be either 'asc' or 'desc'. ")
		}
	} else {
		sort = defaultSort
	}

	return intLimit, intPage, orderby, sort, nil
}
