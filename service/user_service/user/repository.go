package user

import (
	"github.com/syedomair/backend-example/models"
)

// Repository interface
type Repository interface {
	GetAllUserDB(limit int, offset int, orderby string, sort string) ([]*models.User, string, error)
	GetUserHighAge() (int, error)
	GetUserLowAge() (int, error)
	GetUserAvgAge() (float64, error)
	GetUserLowSalary() (float64, error)
	GetUserHighSalary() (float64, error)
	GetUserAvgSalary() (float64, error)
}
