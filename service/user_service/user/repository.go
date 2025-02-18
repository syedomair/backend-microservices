package user

import "github.com/syedomair/backend-example/models"

// Repository interface
type Repository interface {
	GetAllUserDB(limit int, offset int, orderby string, sort string) ([]*models.User, string, error)
}
