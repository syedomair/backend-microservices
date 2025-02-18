package models

// Department Public
type Department struct {
	ID      string `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

// TableName Public
func (Department) TableName() string {
	return "department"
}
