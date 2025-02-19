package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/syedomair/backend-example/lib/request"
	"github.com/syedomair/backend-example/lib/response"
	"github.com/syedomair/backend-example/models"
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

	responseObj, err := c.getAllUsersData(limit, offset, orderBy, sort)
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

// getAllUsersData fetches user data and statistics concurrently.
func (c *Controller) getAllUsersData(limit, offset int, orderBy, sort string) (map[string]interface{}, error) {
	methodName := "GetAllUsers"
	c.Logger.Debug("method start", zap.String("method", methodName))
	start := time.Now()
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

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

	responseObj := map[string]interface{}{
		"high_age":    strconv.Itoa(intUserHighAge),
		"low_age":     strconv.Itoa(intUserLowAge),
		"avg_age":     fmt.Sprintf("%.2f", fltUserAvgAge),
		"high_salary": fmt.Sprintf("%.2f", fltUserHighSalary),
		"low_salary":  fmt.Sprintf("%.2f", fltUserLowSalary),
		"avg_salary":  fmt.Sprintf("%.2f", fltUserAvgSalary),
		"count":       count,
		"list":        userList,
	}

	c.Logger.Debug("method end", zap.String("method", methodName), zap.Duration("duration", time.Since(start)))
	return responseObj, nil

}
