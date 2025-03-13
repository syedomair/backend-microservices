package point

// Repository interface
type Repository interface {
	GetUserPointDB(userID string) (int, error)
}
