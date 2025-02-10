package user

import (
	"backend/models"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type postgresRepo struct {
	client    *gorm.DB
	logger    *zap.Logger
	requestID string
}

// NewPostgresRepository Public.
func NewPostgresRepository(c *gorm.DB, logger *zap.Logger) Repository {
	return &postgresRepo{client: c, logger: logger}
}

func (p *postgresRepo) SetRequestID(requestID string) {
	p.requestID = requestID
}

// GetAllActionsDB Public
func (p *postgresRepo) GetAllActionDB(limit int, offset int, orderby string, sort string) ([]*models.Action, string, error) {
	methodName := "GetAllActionsDB"
	p.logger.Debug("GetAllActionDB", zap.String("m start", methodName))
	start := time.Now()
	actions := []*models.Action{}
	count := int64(0)
	if err := p.client.Table("action").
		Select("*").
		Limit(limit).
		Offset(offset).
		Order("action.taskid, action.nexttaskid").
		Count(&count).
		Scan(&actions).Error; err != nil {
		return nil, "", err
	}
	p.logger.Debug(p.requestID, zap.String("m", methodName), zap.Duration("since", time.Since(start)))

	return actions, "", nil
}
