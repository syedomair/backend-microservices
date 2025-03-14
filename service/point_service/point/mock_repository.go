package point

// MockRepository is a manual mock implementation of the Repository interface.
type MockRepositoryDB struct {
	GetUserPointDBFunc func(userID string) (int, error)
}

func (m *MockRepositoryDB) GetUserPointDB(userID string) (int, error) {
	return m.GetUserPointDBFunc(userID)
}
