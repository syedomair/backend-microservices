package department

import (
	"github.com/syedomair/backend-microservices/models"
)

// Repository interface
type Repository interface {
	GetAllDepartmentDB(limit int, offset int, orderby string, sort string) ([]*models.Department, string, error)
}
