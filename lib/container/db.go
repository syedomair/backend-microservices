package container

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db interface {
	MakeConnection() (*gorm.DB, error)
}

func NewDBConnectionAdapter(dbName, url string, dbMaxIdle, dbMaxOpen, dbMaxLifeTime, dbMaxIdleTime int, gormConf string) Db {
	switch dbName {
	case Postgres:
		return &PostgresAdapter{dbUrl: url,
			dbMaxIdle:     dbMaxIdle,
			dbMaxOpen:     dbMaxOpen,
			dbMaxLifeTime: dbMaxLifeTime,
			dbMaxIdleTime: dbMaxIdleTime,
			gormConf:      gormConf}
	case Mysql:
		return &MySQLAdapter{dbUrl: url,
			dbMaxIdle:     dbMaxIdle,
			dbMaxOpen:     dbMaxOpen,
			dbMaxLifeTime: dbMaxLifeTime,
			dbMaxIdleTime: dbMaxIdleTime,
			gormConf:      gormConf}
	}
	return &PostgresAdapter{dbUrl: url,
		dbMaxIdle:     dbMaxIdle,
		dbMaxOpen:     dbMaxOpen,
		dbMaxLifeTime: dbMaxLifeTime,
		dbMaxIdleTime: dbMaxIdleTime,
		gormConf:      gormConf}
}

type PostgresAdapter struct {
	dbUrl         string
	dbMaxIdle     int
	dbMaxOpen     int
	dbMaxLifeTime int
	dbMaxIdleTime int
	gormConf      string
}

var _ Db = (*PostgresAdapter)(nil)

func NewPostgresAdapter(url string, dbMaxIdle, dbMaxOpen, dbMaxLifeTime, dbMaxIdleTime int, gormConf string) *PostgresAdapter {
	return &PostgresAdapter{dbUrl: url,
		dbMaxIdle:     dbMaxIdle,
		dbMaxOpen:     dbMaxOpen,
		dbMaxLifeTime: dbMaxLifeTime,
		dbMaxIdleTime: dbMaxIdleTime,
		gormConf:      gormConf,
	}
}

func (p *PostgresAdapter) MakeConnection() (*gorm.DB, error) {
	return makeConnection(postgres.Open(p.dbUrl), p.dbUrl, p.dbMaxIdle, p.dbMaxOpen, p.dbMaxLifeTime, p.dbMaxIdleTime, p.gormConf)
}

type MySQLAdapter struct {
	dbUrl         string
	dbMaxIdle     int
	dbMaxOpen     int
	dbMaxLifeTime int
	dbMaxIdleTime int
	gormConf      string
}

var _ Db = (*MySQLAdapter)(nil)

func NewMySQLAdapter(url string, dbMaxIdle, dbMaxOpen, dbMaxLifeTime, dbMaxIdleTime int, gormConf string) *MySQLAdapter {
	return &MySQLAdapter{dbUrl: url,
		dbMaxIdle:     dbMaxIdle,
		dbMaxOpen:     dbMaxOpen,
		dbMaxLifeTime: dbMaxLifeTime,
		dbMaxIdleTime: dbMaxIdleTime,
		gormConf:      gormConf,
	}
}

func (m *MySQLAdapter) MakeConnection() (*gorm.DB, error) {
	return makeConnection(mysql.Open(m.dbUrl), m.dbUrl, m.dbMaxIdle, m.dbMaxOpen, m.dbMaxLifeTime, m.dbMaxIdleTime, m.gormConf)
}

func makeConnection(dialector gorm.Dialector, url string, dbMaxIdle, dbMaxOpen, dbMaxLifeTime, dbMaxIdleTime int, gormConf string) (*gorm.DB, error) {
	if url == "" {
		return nil, errors.New("database URL is required")
	}

	file, err := os.Open(gormConf)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer file.Close()

	type tempLoggerConfig struct {
		SlowThreshold             int  `json:"slow_threshold"`
		Colorful                  bool `json:"colorful"`
		IgnoreRecordNotFoundError bool `json:"ignore_record_not_found_error"`
		ParameterizedQueries      bool `json:"parameterized_queries"`
		LogLevel                  int  `json:"Log_level"`
	}

	var cfg tempLoggerConfig
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(cfg.SlowThreshold) * time.Second,
			Colorful:                  cfg.Colorful,
			IgnoreRecordNotFoundError: cfg.IgnoreRecordNotFoundError,
			ParameterizedQueries:      cfg.ParameterizedQueries,
			LogLevel:                  logger.LogLevel(cfg.LogLevel),
		},
	)

	gormDB, err := gorm.Open(dialector, &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(dbMaxIdle)
	sqlDB.SetMaxOpenConns(dbMaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifeTime) * time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbMaxIdleTime) * time.Minute)

	return gormDB, nil
}
