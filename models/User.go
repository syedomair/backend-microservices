package models

// User Public
type User struct {
	ID           string  `json:"id" gorm:"column:id"`
	Name         string  `json:"name" gorm:"column:name"`
	Email        string  `json:"email" gorm:"column:email"`
	DepartmentID string  `json:"department_id" gorm:"column:department_id"`
	Age          int     `json:"age" gorm:"column:age"`
	Salary       float32 `json:"salary" gorm:"column:salary"`
}

// TableName Public
func (User) TableName() string {
	return "public.user"
}
