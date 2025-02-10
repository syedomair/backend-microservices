package user

import "backend/models"

// Repository interface
type Repository interface {
	SetRequestID(requestID string)
	GetAllActionDB(limit int, offset int, orderby string, sort string) ([]*models.Action, string, error)
}
