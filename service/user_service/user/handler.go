package user

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/syedomair/backend-microservices/lib/request"
	"github.com/syedomair/backend-microservices/lib/response"
	"github.com/syedomair/backend-microservices/models"

	"go.uber.org/zap"
)

type Controller struct {
	Logger *zap.Logger
	Repo   Repository
}

// GetAllUsers retrieves all users with additional statistics.
func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	methodName := "GetAllUsers"
	c.Logger.Debug("method start", zap.String("method", methodName))
	start := time.Now()

	limit, offset, orderBy, sort, err := request.ValidateQueryString(r, "1000", "0", "name", "asc")
	if err != nil {
		c.handleError(methodName, w, err, http.StatusBadRequest)
		return
	}

	userStatistics, err := c.GetUserStatistics(limit, offset, orderBy, sort)
	if err != nil {
		c.handleError(methodName, w, err, http.StatusBadRequest)
		return
	}

	responseUserObj := models.ResponseUser{
		HighAge:    strconv.Itoa(userStatistics.UserHighAge),
		LowAge:     strconv.Itoa(userStatistics.UserLowAge),
		AvgAge:     fmt.Sprintf("%.2f", userStatistics.UserAvgAge),
		HighSalary: fmt.Sprintf("%.2f", userStatistics.UserHighSalary),
		LowSalary:  fmt.Sprintf("%.2f", userStatistics.UserLowSalary),
		AvgSalary:  fmt.Sprintf("%.2f", userStatistics.UserAvgSalary),
		Count:      userStatistics.Count,
		List:       userStatistics.UserList,
	}

	var responseObj map[string]interface{}
	err = mapstructure.Decode(responseUserObj, &responseObj)
	if err != nil {
		c.handleError(methodName, w, err, http.StatusBadRequest)
		return
	}

	c.Logger.Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	response.SuccessResponseHelper(w, responseObj, http.StatusOK)
}
func (c *Controller) GetUserStatistics(limit, offset int, orderBy, sort string) (*models.UserStatistics, error) {
	methodName := "GetAllUsers"
	c.Logger.Debug("method start", zap.String("method", methodName))
	start := time.Now()

	userService := NewUserService(c.Repo, c.Logger)

	userStatistics, err := userService.GetAllUserStatistics(limit, offset, orderBy, sort)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return userStatistics, nil
}

func (c *Controller) handleError(methodName string, w http.ResponseWriter, err error, statusCode int) {
	c.Logger.Error("method failed", zap.String("method", methodName), zap.Error(err))
	response.ErrorResponseHelper(methodName, w, err.Error(), statusCode)
}
