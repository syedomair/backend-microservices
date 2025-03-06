package user

import (
	"context"
	"time"

	"github.com/syedomair/backend-microservices/models"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type UserService struct {
	repo   Repository
	logger *zap.Logger
}

func NewUserService(repo Repository, logger *zap.Logger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

// GetAllUserStatistics
func (u *UserService) GetAllUserStatistics(limit, offset int, orderBy, sort string) (*models.UserStatistics, error) {
	methodName := "GetAllUserStatistics"
	u.logger.Debug("method start", zap.String("method_name", methodName))
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
		userList, count, err = u.repo.GetAllUserDB(limit, offset, orderBy, sort)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		intUserHighAge, err = u.repo.GetUserHighAge()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		intUserLowAge, err = u.repo.GetUserLowAge()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserAvgAge, err = u.repo.GetUserAvgAge()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserLowSalary, err = u.repo.GetUserLowSalary()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserHighSalary, err = u.repo.GetUserHighSalary()
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		var err error
		fltUserAvgSalary, err = u.repo.GetUserAvgSalary()
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	userStatistics := &models.UserStatistics{
		UserList:       userList,
		Count:          count,
		UserHighAge:    intUserHighAge,
		UserLowAge:     intUserLowAge,
		UserAvgAge:     fltUserAvgAge,
		UserLowSalary:  fltUserLowSalary,
		UserHighSalary: fltUserHighSalary,
		UserAvgSalary:  fltUserAvgSalary,
	}

	u.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return userStatistics, nil
}
