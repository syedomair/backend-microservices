package pointserver

// Repository interface
type Repository interface {
	GetUserPointDB(userID string) (int, error)
}
