package point

import (
	"time"

	"go.uber.org/zap"
)

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
		return points, err
	}

	p.logger.Debug("method end", zap.String("method_name", methodName), zap.Duration("since", time.Since(start)))
	return points, nil
}
