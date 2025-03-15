package point

// Repository interface
type Repository interface {
	GetUserPointDB(userID string) (int, error)
	GetUserListPointsDB(userIDs []string) (map[string]int32, error)
}
