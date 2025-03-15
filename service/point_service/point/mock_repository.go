package point

// MockRepository is a manual mock implementation of the Repository interface.
type MockRepositoryDB struct {
	GetUserPointDBFunc      func(userID string) (int, error)
	GetUserListPointsDBFunc func(userIDs []string) (map[string]int32, error)
}

func (m *MockRepositoryDB) GetUserPointDB(userID string) (int, error) {
	return m.GetUserPointDBFunc(userID)
}
func (m *MockRepositoryDB) GetUserListPointsDB(userIDs []string) (map[string]int32, error) {
	return m.GetUserListPointsDBFunc(userIDs)
}
