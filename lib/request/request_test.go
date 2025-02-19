package request

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestValidateQueryString(t *testing.T) {
	tests := []struct {
		name           string
		query          url.Values
		defaultLimit   string
		defaultPage    string
		defaultOrderby string
		defaultSort    string
		wantLimit      int
		wantPage       int
		wantOrderby    string
		wantSort       string
		wantErr        bool
	}{
		{
			name:           "Valid query parameters",
			query:          url.Values{"limit": {"10"}, "page": {"2"}, "orderby": {"name"}, "sort": {"asc"}},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      10,
			wantPage:       10,
			wantOrderby:    "name",
			wantSort:       "asc",
			wantErr:        false,
		},
		{
			name:           "Default values",
			query:          url.Values{},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      20,
			wantPage:       0,
			wantOrderby:    "id",
			wantSort:       "desc",
			wantErr:        false,
		},
		{
			name:           "Invalid limit",
			query:          url.Values{"limit": {"abc"}, "page": {"1"}, "orderby": {"name"}, "sort": {"asc"}},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      0,
			wantPage:       0,
			wantOrderby:    "",
			wantSort:       "",
			wantErr:        true,
		},
		{
			name:           "Invalid page",
			query:          url.Values{"limit": {"10"}, "page": {"abc"}, "orderby": {"name"}, "sort": {"asc"}},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      0,
			wantPage:       0,
			wantOrderby:    "",
			wantSort:       "",
			wantErr:        true,
		},
		{
			name:           "Invalid orderby",
			query:          url.Values{"limit": {"10"}, "page": {"1"}, "orderby": {"123"}, "sort": {"asc"}},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      0,
			wantPage:       0,
			wantOrderby:    "",
			wantSort:       "",
			wantErr:        true,
		},
		{
			name:           "Invalid sort",
			query:          url.Values{"limit": {"10"}, "page": {"1"}, "orderby": {"name"}, "sort": {"123"}},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      0,
			wantPage:       0,
			wantOrderby:    "",
			wantSort:       "",
			wantErr:        true,
		},
		{
			name:           "Invalid sort value",
			query:          url.Values{"limit": {"10"}, "page": {"1"}, "orderby": {"name"}, "sort": {"invalid"}},
			defaultLimit:   "20",
			defaultPage:    "1",
			defaultOrderby: "id",
			defaultSort:    "desc",
			wantLimit:      0,
			wantPage:       0,
			wantOrderby:    "",
			wantSort:       "",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/?"+tt.query.Encode(), nil)
			gotLimit, gotPage, gotOrderby, gotSort, err := ValidateQueryString(req, tt.defaultLimit, tt.defaultPage, tt.defaultOrderby, tt.defaultSort)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateQueryString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("ValidateQueryString() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
			if gotPage != tt.wantPage {
				t.Errorf("ValidateQueryString() gotPage = %v, want %v", gotPage, tt.wantPage)
			}
			if gotOrderby != tt.wantOrderby {
				t.Errorf("ValidateQueryString() gotOrderby = %v, want %v", gotOrderby, tt.wantOrderby)
			}
			if gotSort != tt.wantSort {
				t.Errorf("ValidateQueryString() gotSort = %v, want %v", gotSort, tt.wantSort)
			}
		})
	}
}
