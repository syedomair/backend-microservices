package point

import (
	"time"

	"go.uber.org/zap"
)

type PointServiceInterface interface {
	GetUserPoints(userID string) (int, error)
	GetUserListPoints(userIDs []string) (map[string]int32, error)
}

type PointService struct {
	repo   Repository
	logger *zap.Logger
}

func NewPointService(repo Repository, logger *zap.Logger) *PointService {
	return &PointService{repo: repo, logger: logger}
}

// GetUserPoints
func (p *PointService) GetUserPoints(userID string) (int, error) {
	methodName := "GetUserPoints"
	p.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	points, err := p.repo.GetUserPointDB(userID)
	if err != nil {
		return 0, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return points, nil
}

// GetUserListPoints
func (m *PointService) GetUserListPoints(userIDs []string) (map[string]int32, error) {
	methodName := "GetUserListPoints"
	m.logger.Debug("method start", zap.String("method_name", methodName))
	start := time.Now()

	points, err := m.repo.GetUserListPointsDB(userIDs)
	if err != nil {
		return nil, err
	}

	m.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return points, nil
}
