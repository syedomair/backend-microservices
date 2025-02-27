package user

import (
	"fmt"
	"testing"

	"github.com/syedomair/backend-microservices/models"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestController_GetAllUsersData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockLogger := zap.NewNop()

	controller := &Controller{
		Repo:   mockRepo,
		Logger: mockLogger,
	}

	tests := []struct {
		name           string
		limit          int
		offset         int
		orderBy        string
		sort           string
		mockSetup      func(*MockRepository)
		expectedResult map[string]interface{}
		expectedError  error
	}{
		{
			name:    "Success - All repository calls succeed",
			limit:   10,
			offset:  0,
			orderBy: "id",
			sort:    "asc",
			mockSetup: func(mr *MockRepository) {
				// Mock all repository calls to return valid data
				mr.EXPECT().GetAllUserDB(10, 0, "id", "asc").Return([]*models.User{{ID: "1", Name: "John"}}, "1", nil).AnyTimes()
				mr.EXPECT().GetUserHighAge().Return(30, nil).AnyTimes() // gomock.Any() matches any context
				mr.EXPECT().GetUserLowAge().Return(20, nil).AnyTimes()
				mr.EXPECT().GetUserAvgAge().Return(25.0, nil).AnyTimes()
				mr.EXPECT().GetUserHighSalary().Return(100000.0, nil).AnyTimes()
				mr.EXPECT().GetUserLowSalary().Return(50000.0, nil).AnyTimes()
				mr.EXPECT().GetUserAvgSalary().Return(75000.0, nil).AnyTimes()
			},
			expectedResult: map[string]interface{}{
				"HighAge":    "30",
				"LowAge":     "20",
				"AvgAge":     "25.00",
				"HighSalary": "100000.00",
				"LowSalary":  "50000.00",
				"AvgSalary":  "75000.00",
				"Count":      "1",
				"List":       []*models.User{{ID: "1", Name: "John"}},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock expectations
			tt.mockSetup(mockRepo)

			// Call the method under test
			result, err := controller.GetAllUsersData(tt.limit, tt.offset, tt.orderBy, tt.sort)

			// Verify the error
			if tt.expectedError != nil {
				if err == nil || err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

			// Verify the result
			if tt.expectedResult != nil {
				if result == nil {
					t.Errorf("expected result: %v, got: %v", tt.expectedResult, result)
				} else {
					for key, expectedValue := range tt.expectedResult {
						if key != "List" {
							actualValue := result[key]
							if fmt.Sprintf("%v", actualValue) != fmt.Sprintf("%v", expectedValue) {
								t.Errorf("expected %s: %v, got: %v", key, expectedValue, actualValue)
							}
						}
					}
				}
			}
		})
	}
}
