package department

import (
	"strconv"
	"time"

	"github.com/syedomair/backend-microservices/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type dbRepo struct {
	client *gorm.DB
	logger *zap.Logger
}

// NewDBRepository Public.
func NewDBRepository(c *gorm.DB, logger *zap.Logger) Repository {
	return &dbRepo{client: c, logger: logger}
}

// GetAllDepartmentDB Public
func (p *dbRepo) GetAllDepartmentDB(limit int, offset int, orderby string, sort string) ([]*models.Department, string, error) {
	methodName := "GetAllDepartmentDB"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	departments := []*models.Department{}
	count := int64(0)
	if err := p.client.Table("department").
		Select("*").
		Limit(limit).
		Offset(offset).
		Order(orderby).
		Scan(&departments).Count(&count).Error; err != nil {
		return nil, "", err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return departments, strconv.Itoa(int(count)), nil
}
