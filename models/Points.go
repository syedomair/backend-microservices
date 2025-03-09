package models

// Points Public
type Points struct {
	ID     string `json:"id" gorm:"column:id"`
	UserID string `json:"user_id" gorm:"column:user_id"`
	Points int    `json:"points" gorm:"column:points"`
}

// TableName Public
func (Points) TableName() string {
	return "points"
}
