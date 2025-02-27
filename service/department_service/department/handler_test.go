package department

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/models"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestGetAllDepartments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	controller := &Controller{
		Logger: logger,
		Repo:   mockRepo,
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		mockReturn     []*models.Department
		mockCount      string
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success",
			queryParams: map[string]string{
				"limit":   "10",
				"offset":  "0",
				"orderby": "name",
				"sort":    "asc",
			},
			mockReturn: []*models.Department{
				{ID: "1", Name: "HR", Address: "123 Main St"},
				{ID: "2", Name: "IT", Address: "456 Tech St"},
			},
			mockCount:      "2",
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Query Parameters",
			queryParams: map[string]string{
				"limit":   "invalid",
				"offset":  "0",
				"orderby": "name",
				"sort":    "asc",
			},
			mockReturn:     nil,
			mockCount:      "",
			mockError:      errors.New("invalid query parameters"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/departments", nil)
			assert.NoError(t, err)

			q := req.URL.Query()
			for key, value := range tt.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()

			if tt.mockError == nil {
				mockRepo.EXPECT().GetAllDepartmentDB(10, 0, "name", "asc").Return(tt.mockReturn, tt.mockCount, tt.mockError)
			}

			controller.GetAllDepartments(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "success", response["result"])
				assert.NotNil(t, response["data"])
			} else {
				var response map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "failure", response["result"])
				assert.NotNil(t, response["data"])
			}
		})
	}
}
