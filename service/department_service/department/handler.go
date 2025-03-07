package department

import (
	"net/http"
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

// GetAllDepartments retrieves all departments with additional statistics.
func (c *Controller) GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	methodName := "GetAllDepartments"
	c.Logger.Debug("method start", zap.String("method", methodName))
	start := time.Now()

	queryParam, err := request.ValidateQueryString(r, "1000", "0", "name", "asc")
	if err != nil {
		c.handleError(methodName, w, err, http.StatusBadRequest)
		return
	}

	responseObj, err := c.GetAllDepartmentData(queryParam.Limit, queryParam.Page, queryParam.OrderBy, queryParam.Sort)
	if err != nil {
		c.handleError(methodName, w, err, http.StatusBadRequest)
		return
	}

	c.Logger.Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	response.SuccessResponseHelper(w, responseObj, http.StatusOK)
}

// handleError abstracts error handling logic.
func (c *Controller) handleError(methodName string, w http.ResponseWriter, err error, statusCode int) {
	c.Logger.Error("method failed", zap.String("method", methodName), zap.Error(err))
	response.ErrorResponseHelper(methodName, w, err.Error(), statusCode)
}

// GetAllDepartmentData fetches department data and statistics concurrently.
func (c *Controller) GetAllDepartmentData(limit, offset int, orderBy, sort string) (map[string]interface{}, error) {
	methodName := "GetAllDepartmentData"
	c.Logger.Debug("method start", zap.String("method", methodName))
	start := time.Now()

	departmentList, count, err := c.Repo.GetAllDepartmentDB(limit, offset, orderBy, sort)
	if err != nil {
		return nil, err
	}

	responseDepartmentObj := models.ResponseDepartment{
		Count: count,
		List:  departmentList,
	}

	var responseObj map[string]interface{}
	err = mapstructure.Decode(responseDepartmentObj, &responseObj)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return responseObj, nil
}
