package user

import (
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
		Count(&count).
		Scan(&users).Error; err != nil {
		return nil, "", err
	}
	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return users, "", nil
}
