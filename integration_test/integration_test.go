package integration_test

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/syedomair/backend-microservices/lib/container"
	"github.com/syedomair/backend-microservices/lib/mockgrpc"
	"github.com/syedomair/backend-microservices/service/department_service/department"
	"github.com/syedomair/backend-microservices/service/user_service/user"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
		container.LogLevel:      "DEBUG",
		container.DatabaseURL:   connecStr,
		container.Port:          "8186",
		container.DBMaxIdle:     "10",
		container.DBMaxOpen:     "100",
		container.DBMaxLifeTime: "1",
		container.DBMaxIdleTime: "10",
		container.ZapConf:       "../config/zap-logger-config.json",
		container.GormConf:      "../config/gorm-logger-config.json",
		container.PprofEnable:   "false",
		container.PointSrvcAddr: "point_service",
		container.PointSrvcMax:  "10",
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

func TestDepartmentAPI(t *testing.T) {
	c := setupTestDB(t)
	departmentRepo := department.NewDBRepository(c.Db(), c.Logger())

	limit := 10
	offset := 0
	orderby := "name"
	sort := "asc"

	departmentDB, count, err := departmentRepo.GetAllDepartmentDB(limit, offset, orderby, sort)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 3, len(departmentDB))
	assert.Equal(t, "3", count)

	mockRepo := departmentRepo
	mockLogger := zap.NewNop()

	controller := &department.Controller{
		Repo:   mockRepo,
		Logger: mockLogger,
	}

	req, err := http.NewRequest("GET", "/departments", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllDepartments(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NoError(t, err)
	assert.Equal(t, "application/json;charset=utf-8", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["result"])
	data := response["data"].(map[string]interface{})
	departments := data["List"].([]interface{})
	department := departments[0].(map[string]interface{})
	assert.Equal(t, "Finance", department["name"])

}

func TestUserAPI(t *testing.T) {
	c := setupTestDB(t)
	userRepo := user.NewDBRepository(c.Db(), c.Logger())

	limit := 10
	offset := 0
	orderby := "name"
	sort := "asc"

	usersDB, count, err := userRepo.GetAllUserDB(limit, offset, orderby, sort)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 9, len(usersDB))
	assert.Equal(t, "9", count)

	mockRepo := userRepo
	mockLogger := zap.NewNop()

	_, conn, _ := mockgrpc.SetupGRPCServer(t)
	defer conn.Close()

	mockConnectionPool := &mockgrpc.MockConnectionPool{
		GetFunc: func() (*grpc.ClientConn, error) {
			return conn, nil
		},
		PutFunc: func(conn *grpc.ClientConn) {
		},
	}

	controller := &user.Controller{
		Repo:                       mockRepo,
		Logger:                     mockLogger,
		PointServiceConnectionPool: mockConnectionPool,
	}

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// Act
	controller.GetAllUsers(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NoError(t, err)
	assert.Equal(t, "application/json;charset=utf-8", rr.Header().Get("Content-Type"))

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["result"])

	data := response["data"].(map[string]interface{})
	assert.Equal(t, "40", data["HighAge"])
	assert.Equal(t, "22", data["LowAge"])
	assert.Equal(t, "31.22", data["AvgAge"])
	assert.Equal(t, "90000.00", data["HighSalary"])
	assert.Equal(t, "48000.00", data["LowSalary"])
	assert.Equal(t, "68333.33", data["AvgSalary"])
	assert.Equal(t, "9", data["Count"])

	users := data["List"].([]interface{})
	assert.Equal(t, 9, len(users))

	user := users[0].(map[string]interface{})
	assert.Equal(t, "Alice Johnson", user["name"])
	assert.Equal(t, float64(30), user["age"])
	assert.Equal(t, 60000.0, user["salary"])
	assert.Equal(t, float64(10), user["point"])
}
