package integration_test

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/lib/container"
	"github.com/syedomair/backend-microservices/service/user_service/user"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "postgres",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections"),
			wait.ForListeningPort("5432/tcp"),
		),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal("Failed to start container:", err)
	}
	defer postgresContainer.Terminate(ctx)

	host, _ := postgresContainer.Host(ctx)
	port, _ := postgresContainer.MappedPort(ctx, "5432")

	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port.Port())
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")

	exitCode := m.Run()
	os.Exit(exitCode)
}

func setupTestDB(t *testing.T) container.Container {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"postgres",
	)

	// Create a unique test database
	dbName := "testdb_" + strings.ToLower(randString(6))
	rootDB, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}
	defer rootDB.Close()

	_, err = rootDB.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		t.Fatal("Failed to create test database:", err)
	}

	// Connect to the new test database
	testConnStr := fmt.Sprintf(
		"%s dbname=%s",
		connStr,
		dbName,
	)
	testDB, err := sql.Open("postgres", testConnStr)
	if err != nil {
		t.Fatal(err)
	}

	// Run migrations
	if err := runMigrations(testDB); err != nil {
		t.Fatal("Failed to run migrations:", err)
	}

	connecStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), dbName)

	c, err := container.New(map[string]string{
		container.LogLevelEnvVar:      "DEBUG",
		container.DatabaseURLEnvVar:   connecStr,
		container.PortEnvVar:          "8186",
		container.DBMaxIdleEnvVar:     "10",
		container.DBMaxOpenEnvVar:     "100",
		container.DBMaxLifeTimeEnvVar: "1",
		container.DBMaxIdleTimeEnvVar: "10",
		container.ZapConf:             "../config/zap-logger-config.json",
		container.GormConf:            "../config/gorm-logger-config.json",
		container.PprofEnable:         "false",
	})
	if err != nil {
		defer func() {
			fmt.Println("server initialization failed error: %w", err)
		}()
		panic("server initialization failed")
	}

	defer testDB.Close()

	return c
}

func runMigrations(db *sql.DB) error {
	// Read and execute your schema.sql
	schema, err := os.ReadFile("../database/migration.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(schema))
	return err
}

func randString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:n]
}

func TestUserRepo(t *testing.T) {
	c := setupTestDB(t)

	userRepo := user.NewPostgresRepository(c.Db(), c.Logger())

	limit := 10
	offset := 0
	orderby := "name"
	sort := "asc"

	users, count, err := userRepo.GetAllUserDB(limit, offset, orderby, sort)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 9, len(users))
	assert.Equal(t, "9", count)

	mockRepo := userRepo
	mockLogger := zap.NewNop()

	controller := &user.Controller{
		Repo:   mockRepo,
		Logger: mockLogger,
	}

	result, err := controller.GetAllUsersData(10, 0, "name", "asc")

	assert.NoError(t, err)
	assert.Equal(t, 8, len(result))

	assert.Equal(t, "40", result["HighAge"].(string))
	assert.Equal(t, "22", result["LowAge"].(string))
	assert.Equal(t, "90000.00", result["HighSalary"].(string))
	assert.Equal(t, "48000.00", result["LowSalary"].(string))
	assert.Equal(t, "31.22", result["AvgAge"].(string))
	assert.Equal(t, "68333.33", result["AvgSalary"].(string))
	assert.Equal(t, "9", result["Count"].(string))

}
