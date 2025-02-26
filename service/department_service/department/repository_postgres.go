package user

import (
	"strconv"
	"time"

	"github.com/syedomair/backend-example/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type postgresRepo struct {
	client *gorm.DB
	logger *zap.Logger
}

// NewPostgresRepository Public.
func NewPostgresRepository(c *gorm.DB, logger *zap.Logger) Repository {
	return &postgresRepo{client: c, logger: logger}
}

// GetAllUserDB Public
func (p *postgresRepo) GetAllUserDB(limit int, offset int, orderby string, sort string) ([]*models.User, string, error) {
	methodName := "GetAllUserDB"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	users := []*models.User{}
	count := int64(0)
	if err := p.client.Table("public.user").
		Select("*").
		Limit(limit).
		Offset(offset).
		Order(orderby).
		Scan(&users).Count(&count).Error; err != nil {
		return nil, "", err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return users, strconv.Itoa(int(count)), nil
}

// GetUserHighAge Public
func (p *postgresRepo) GetUserHighAge() (int, error) {
	methodName := "GetUserHighAge"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	var highAge int
	if err := p.client.Table("public.user").
		Select("MAX(age)").
		Scan(&highAge).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return highAge, nil
}

// GetUserLowAge Public
func (p *postgresRepo) GetUserLowAge() (int, error) {
	methodName := "GetUserLowAge"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	var lowAge int
	if err := p.client.Table("public.user").
		Select("MIN(age)").
		Scan(&lowAge).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return lowAge, nil
}

// GetUserAvgAge Public
func (p *postgresRepo) GetUserAvgAge() (float64, error) {
	methodName := "GetUserAvgAge"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	var avgAge float64
	if err := p.client.Table("public.user").
		Select("AVG(age)").
		Scan(&avgAge).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return avgAge, nil
}

// GetUserLowSalary Public
func (p *postgresRepo) GetUserLowSalary() (float64, error) {
	methodName := "GetUserLowSalary"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	var lowSalary float64
	if err := p.client.Table("public.user").
		Select("MIN(salary)").
		Scan(&lowSalary).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return lowSalary, nil
}

// GetUserHighSalary Public
func (p *postgresRepo) GetUserHighSalary() (float64, error) {
	methodName := "GetUserHighSalary"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	var highSalary float64
	if err := p.client.Table("public.user").
		Select("MAX(salary)").
		Scan(&highSalary).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return highSalary, nil
}

// GetUserAvgSalary Public
func (p *postgresRepo) GetUserAvgSalary() (float64, error) {
	methodName := "GetUserAvgSalary"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	var avgSalary float64
	if err := p.client.Table("public.user").
		Select("AVG(salary)").
		Scan(&avgSalary).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return avgSalary, nil
}
