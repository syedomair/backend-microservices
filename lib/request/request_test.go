package request

import (
	"net/http"
	"testing"
)

func TestValidateQueryString(t *testing.T) {
	tests := []struct {
		name           string
		queryString    string
		defaultLimit   string
		defaultPage    string
		defaultOrderby string
		defaultSort    string
		want           QueryParams
		wantErr        bool
	}{
		{
			name:           "default values",
			queryString:    "",
			defaultLimit:   "10",
			defaultPage:    "1",
			defaultOrderby: "name",
			defaultSort:    "asc",
			want: QueryParams{
				Limit:   10,
				Page:    0,
				OrderBy: "name",
				Sort:    "asc",
			},
			wantErr: false,
		},
		{
			name:           "custom values",
			queryString:    "limit=20&page=2&orderby=id&sort=desc",
			defaultLimit:   "10",
			defaultPage:    "1",
			defaultOrderby: "name",
			defaultSort:    "asc",
			want: QueryParams{
				Limit:   20,
				Page:    20,
				OrderBy: "id",
				Sort:    "desc",
			},
			wantErr: false,
		},
		{
			name:           "invalid limit",
			queryString:    "limit=abc",
			defaultLimit:   "10",
			defaultPage:    "1",
			defaultOrderby: "name",
			defaultSort:    "asc",
			want:           QueryParams{},
			wantErr:        true,
		},
		{
			name:           "invalid page",
			queryString:    "page=abc",
			defaultLimit:   "10",
			defaultPage:    "1",
			defaultOrderby: "name",
			defaultSort:    "asc",
			want:           QueryParams{},
			wantErr:        true,
		},
		{
			name:           "invalid orderby",
			queryString:    "orderby=123",
			defaultLimit:   "10",
			defaultPage:    "1",
			defaultOrderby: "name",
			defaultSort:    "asc",
			want:           QueryParams{},
			wantErr:        true,
		},
		{
			name:           "invalid sort",
			queryString:    "sort=invalid",
			defaultLimit:   "10",
			defaultPage:    "1",
			defaultOrderby: "name",
			defaultSort:    "asc",
			want:           QueryParams{},
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?"+tt.queryString, nil)
			if err != nil {
				t.Fatal(err)
			}

			got, err := ValidateQueryString(req, tt.defaultLimit, tt.defaultPage, tt.defaultOrderby, tt.defaultSort)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateQueryString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Limit != tt.want.Limit {
					t.Errorf("ValidateQueryString().Limit = %v, want %v", got.Limit, tt.want.Limit)
				}
				if got.Page != tt.want.Page {
					t.Errorf("ValidateQueryString().Page = %v, want %v", got.Page, tt.want.Page)
				}
				if got.OrderBy != tt.want.OrderBy {
					t.Errorf("ValidateQueryString().OrderBy = %v, want %v", got.OrderBy, tt.want.OrderBy)
				}
				if got.Sort != tt.want.Sort {
					t.Errorf("ValidateQueryString().Sort = %v, want %v", got.Sort, tt.want.Sort)
				}
			}
		})
	}
}

func TestGetLimit(t *testing.T) {
	tests := []struct {
		name         string
		queryString  string
		defaultLimit string
		want         int
		wantErr      bool
	}{
		{
			name:         "default value",
			queryString:  "",
			defaultLimit: "10",
			want:         10,
			wantErr:      false,
		},
		{
			name:         "custom value",
			queryString:  "limit=20",
			defaultLimit: "10",
			want:         20,
			wantErr:      false,
		},
		{
			name:         "invalid value",
			queryString:  "limit=abc",
			defaultLimit: "10",
			want:         0,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?"+tt.queryString, nil)
			if err != nil {
				t.Fatal(err)
			}

			got, err := getLimit(req, tt.defaultLimit)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPage(t *testing.T) {
	tests := []struct {
		name        string
		queryString string
		defaultPage string
		want        int
		wantErr     bool
	}{
		{
			name:        "default value",
			queryString: "",
			defaultPage: "1",
			want:        1,
			wantErr:     false,
		},
		{
			name:        "custom value",
			queryString: "page=2",
			defaultPage: "1",
			want:        2,
			wantErr:     false,
		},
		{
			name:        "invalid value",
			queryString: "page=abc",
			defaultPage: "1",
			want:        0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?"+tt.queryString, nil)
			if err != nil {
				t.Fatal(err)
			}

			got, err := getPage(req, tt.defaultPage)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrderBy(t *testing.T) {
	tests := []struct {
		name           string
		queryString    string
		defaultOrderby string
		want           string
		wantErr        bool
	}{
		{
			name:           "default value",
			queryString:    "",
			defaultOrderby: "name",
			want:           "name",
			wantErr:        false,
		},
		{
			name:           "custom value",
			queryString:    "orderby=id",
			defaultOrderby: "name",
			want:           "id",
			wantErr:        false,
		},
		{
			name:           "invalid value",
			queryString:    "orderby=123",
			defaultOrderby: "name",
			want:           "",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?"+tt.queryString, nil)
			if err != nil {
				t.Fatal(err)
			}

			got, err := getOrderBy(req, tt.defaultOrderby)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOrderBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSort(t *testing.T) {
	tests := []struct {
		name        string
		queryString string
		defaultSort string
		want        string
		wantErr     bool
	}{
		{
			name:        "default value",
			queryString: "",
			defaultSort: "asc",
			want:        "asc",
			wantErr:     false,
		},
		{
			name:        "custom value",
			queryString: "sort=desc",
			defaultSort: "asc",
			want:        "desc",
			wantErr:     false,
		},
		{
			name:        "invalid value",
			queryString: "sort=invalid",
			defaultSort: "asc",
			want:        "",
			wantErr:     true,
		},
		{
			name:        "numeric value",
			queryString: "sort=123",
			defaultSort: "asc",
			want:        "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?"+tt.queryString, nil)
			if err != nil {
				t.Fatal(err)
			}

			got, err := getSort(req, tt.defaultSort)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    int
		wantErr bool
	}{
		{
			name:    "valid number",
			str:     "10",
			want:    10,
			wantErr: false,
		},
		{
			name:    "invalid number",
			str:     "abc",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInt(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
