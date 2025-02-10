package db

import (
	"backend/lib/envconstant"
	"errors"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MockDialector is a mock implementation of gorm.Dialector
type MockDialector struct {
	OpenFunc func() (gorm.Dialector, error)
}

func (m *MockDialector) Open() (gorm.Dialector, error) {
	return m.OpenFunc()
}

// TestPostgresAdapter_MakeConnection tests the MakeConnection method of PostgresAdapter
func TestPostgresAdapter_MakeConnection(t *testing.T) {
	tests := []struct {
		name        string
		dbUrl       string
		expectedErr error
	}{
		{
			name:        "Successful connection",
			dbUrl:       os.Getenv(envconstant.DatabaseURLEnvVar),
			expectedErr: nil,
		},
		{
			name:        "Empty URL",
			dbUrl:       "",
			expectedErr: errors.New("database URL is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &PostgresAdapter{
				dbUrl:         tt.dbUrl,
				dbMaxIdle:     10,
				dbMaxOpen:     100,
				dbMaxLifeTime: 1,
				dbMaxIdleTime: 10,
				gormConf:      "../../config/gorm-logger-config.json",
			}

			_, err := adapter.MakeConnection()
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("PostgresAdapter.MakeConnection() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("PostgresAdapter.MakeConnection() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}

// TestMySQLAdapter_MakeConnection tests the MakeConnection method of MySQLAdapter
func TestMySQLAdapter_MakeConnection(t *testing.T) {
	tests := []struct {
		name        string
		dbUrl       string
		expectedErr error
	}{
		/*
			{
				name:        "Successful connection",
				dbUrl:       os.Getenv(envconstant.DatabaseURLEnvVar),
				expectedErr: nil,
			},
		*/
		{
			name:        "Empty URL",
			dbUrl:       "",
			expectedErr: errors.New("database URL is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &MySQLAdapter{
				dbUrl:         tt.dbUrl,
				dbMaxIdle:     10,
				dbMaxOpen:     100,
				dbMaxLifeTime: 1,
				dbMaxIdleTime: 10,
			}

			_, err := adapter.MakeConnection()
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("MySQLAdapter.MakeConnection() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("MySQLAdapter.MakeConnection() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}

// TestNewDBConnectionAdapter tests the NewDBConnectionAdapter function
func TestNewDBConnectionAdapter(t *testing.T) {
	tests := []struct {
		name     string
		dbName   string
		expected Db
	}{
		{
			name:     "Postgres adapter",
			dbName:   envconstant.Postgres,
			expected: &PostgresAdapter{},
		},
		/*
			{
				name:     "MySQL adapter",
				dbName:   envconstant.Mysql,
				expected: &MySQLAdapter{},
			},
		*/
		{
			name:     "Default adapter (Postgres)",
			dbName:   "unknown",
			expected: &PostgresAdapter{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := NewDBConnectionAdapter(tt.dbName, "url", 10, 100, 1, 10, "")
			switch adapter.(type) {
			case *PostgresAdapter:
				if tt.dbName != envconstant.Postgres && tt.dbName != "unknown" {
					t.Errorf("NewDBConnectionAdapter() returned unexpected adapter type for dbName %v", tt.dbName)
				}
			case *MySQLAdapter:
				if tt.dbName != envconstant.Mysql {
					t.Errorf("NewDBConnectionAdapter() returned unexpected adapter type for dbName %v", tt.dbName)
				}
			default:
				t.Errorf("NewDBConnectionAdapter() returned unexpected adapter type")
			}
		})
	}
}

// TestMakeConnection tests the makeConnection function
func TestMakeConnection(t *testing.T) {
	tests := []struct {
		name        string
		dialector   gorm.Dialector
		dbUrl       string
		expectedErr error
	}{
		{
			name:        "Successful connection",
			dialector:   postgres.Open(os.Getenv(envconstant.DatabaseURLEnvVar)),
			dbUrl:       os.Getenv(envconstant.DatabaseURLEnvVar),
			expectedErr: nil,
		},
		{
			name:        "Empty URL",
			dialector:   postgres.Open(""),
			dbUrl:       "",
			expectedErr: errors.New("database URL is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := makeConnection(tt.dialector, tt.dbUrl, 10, 100, 1, 10, "../../config/gorm-logger-config.json")
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("makeConnection() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("makeConnection() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
