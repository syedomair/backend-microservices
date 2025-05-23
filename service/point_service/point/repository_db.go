package point

import (
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

// GetUserPointDB Public
func (p *dbRepo) GetUserPointDB(userID string) (int, error) {
	methodName := "GetUserPointDB"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	points := models.Points{}
	if err := p.client.
		Where("user_id = ?", userID).
		Find(&points).Error; err != nil {
		return 0, nil
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return points.Points, nil
}

// GetUserListPointsDB Public
func (p *dbRepo) GetUserListPointsDB(userIDs []string) (map[string]int32, error) {
	methodName := "GetUserListPointsDB"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	points := []models.Points{}
	if err := p.client.
		Where("user_id in (?)", userIDs).
		Find(&points).Error; err != nil {
		return nil, nil
	}

	mapUserPoints := make(map[string]int32)
	for _, userPoint := range points {
		mapUserPoints[userPoint.UserID] = int32(userPoint.Points)
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return mapUserPoints, nil
}
