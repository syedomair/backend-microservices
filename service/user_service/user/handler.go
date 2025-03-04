package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/syedomair/backend-microservices/lib/request"
	"github.com/syedomair/backend-microservices/lib/response"
	"github.com/syedomair/backend-microservices/models"
	"golang.org/x/sync/errgroup"

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

	responseObj, err := c.GetAllUsersData(limit, offset, orderBy, sort)
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
	response.ErrorResponseHelper(methodName, c.Logger, w, err.Error(), statusCode)
}

// GetAllUsersData fetches user data and statistics concurrently.
func (c *Controller) GetAllUsersData(limit, offset int, orderBy, sort string) (map[string]interface{}, error) {
	methodName := "GetAllUsersData"
	c.Logger.Debug("method start", zap.String("method", methodName))
	start := time.Now()

	g, _ := errgroup.WithContext(context.Background())

	var (
		userList          []*models.User
		count             string
		intUserHighAge    int
		intUserLowAge     int
		fltUserAvgAge     float64
		fltUserAvgSalary  float64
		fltUserLowSalary  float64
		fltUserHighSalary float64
	)

	g.Go(func() error {
		var err error
		userList, count, err = c.Repo.GetAllUserDB(limit, offset, orderBy, sort)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		intUserHighAge, err = c.Repo.GetUserHighAge()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		intUserLowAge, err = c.Repo.GetUserLowAge()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserAvgAge, err = c.Repo.GetUserAvgAge()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserLowSalary, err = c.Repo.GetUserLowSalary()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserHighSalary, err = c.Repo.GetUserHighSalary()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserAvgSalary, err = c.Repo.GetUserAvgSalary()
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}
	responseUserObj := models.ResponseUser{
		HighAge:    strconv.Itoa(intUserHighAge),
		LowAge:     strconv.Itoa(intUserLowAge),
		AvgAge:     fmt.Sprintf("%.2f", fltUserAvgAge),
		HighSalary: fmt.Sprintf("%.2f", fltUserHighSalary),
		LowSalary:  fmt.Sprintf("%.2f", fltUserLowSalary),
		AvgSalary:  fmt.Sprintf("%.2f", fltUserAvgSalary),
		Count:      count,
		List:       userList,
	}

	var responseObj map[string]interface{}
	err := mapstructure.Decode(responseUserObj, &responseObj)
	if err != nil {
		return nil, err
	}

	c.Logger.Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return responseObj, nil

}
