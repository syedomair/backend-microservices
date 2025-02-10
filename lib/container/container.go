package container

import (
	"backend/lib/db"
	"backend/lib/envconstant"
	"backend/lib/logger"
	"fmt"
	"maps"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Container interface it is used to as container
type Container interface {
	Logger() *zap.Logger
	Db() *gorm.DB
	Port() string
	PprofEnable() string
}

type container struct {
	logger               *zap.Logger
	db                   *gorm.DB
	port                 string
	pprofEnable          string
	environmentVariables map[string]string
}

var _ Container = (*container)(nil)

func (c *container) Db() *gorm.DB {
	return c.db
}
func (c *container) dbSetup() (*gorm.DB, error) {

	strDBMaxIdleEnvVar, err := c.getRequiredEnvVar(envconstant.DBMaxIdleEnvVar)
	if err != nil {
		return nil, err
	}
	intDBMaxIdleEnvVar, err := strconv.Atoi(strDBMaxIdleEnvVar)
	if err != nil {
		return nil, err
	}

	strDBMaxOpenEnvVar, err := c.getRequiredEnvVar(envconstant.DBMaxOpenEnvVar)
	if err != nil {
		return nil, err
	}
	intDBMaxOpenEnvVar, err := strconv.Atoi(strDBMaxOpenEnvVar)
	if err != nil {
		return nil, err
	}

	strDBMaxLifeTimeEnvVar, err := c.getRequiredEnvVar(envconstant.DBMaxLifeTimeEnvVar)
	if err != nil {
		return nil, err
	}
	intDBMaxLifeTimeEnvVar, err := strconv.Atoi(strDBMaxLifeTimeEnvVar)
	if err != nil {
		return nil, err
	}

	strDBMaxIdleTimeEnvVar, err := c.getRequiredEnvVar(envconstant.DBMaxIdleTimeEnvVar)
	if err != nil {
		return nil, err
	}
	intDBMaxIdleTimeEnvVar, err := strconv.Atoi(strDBMaxIdleTimeEnvVar)
	if err != nil {
		return nil, err
	}

	strDatabaseURLEnvVar, err := c.getRequiredEnvVar(envconstant.DatabaseURLEnvVar)
	if err != nil {
		return nil, err
	}

	strGormConf, err := c.getRequiredEnvVar(envconstant.GormConf)
	if err != nil {
		return nil, err
	}

	ca := db.NewDBConnectionAdapter(envconstant.DBEnvVar,
		strDatabaseURLEnvVar,
		intDBMaxIdleEnvVar,
		intDBMaxOpenEnvVar,
		intDBMaxLifeTimeEnvVar,
		intDBMaxIdleTimeEnvVar,
		strGormConf)
	if c.db == nil {
		c.db, err = ca.MakeConnection()
		if err != nil {
			return nil, err
		}
	}
	return c.db, nil
}

func (c *container) loggerSetup() (*zap.Logger, error) {
	if c.logger == nil {
		zapConfig, err := c.getRequiredEnvVar(envconstant.ZapConf)
		if err != nil {
			return nil, err
		}
		c.logger, err = logger.New(zapConfig)
		if err != nil {
			return nil, err
		}
	}
	return c.logger, nil
}
func (c *container) Logger() *zap.Logger {
	return c.logger
}

func (c *container) portSetup() (string, error) {
	return c.getRequiredEnvVar(envconstant.PortEnvVar)
}
func (c *container) Port() string {
	return c.port
}

func (c *container) pprofEnableSetup() (string, error) {
	return c.getRequiredEnvVar(envconstant.PprofEnable)
}
func (c *container) PprofEnable() string {
	return c.pprofEnable
}

func (c *container) getRequiredEnvVar(key string) (string, error) {
	value, ok := c.environmentVariables[key]
	if !ok {
		return "", fmt.Errorf("missing mandatory envvar: %q", key)
	}
	return value, nil
}

func New(envVars map[string]string) (Container, error) {
	requiredKeys := maps.Keys(envVars)
	requiredKeysSlice := make([]string, len(envVars))
	for key := range requiredKeys {
		requiredKeysSlice = append(requiredKeysSlice, key)
	}
	if err := validateEnvVars(envVars, requiredKeysSlice); err != nil {
		return &container{environmentVariables: envVars}, err
	}
	c := &container{environmentVariables: envVars}

	var err error
	c.db, err = c.dbSetup()
	if err != nil {
		return c, err
	}
	c.logger, err = c.loggerSetup()
	if err != nil {
		return c, err
	}
	c.port, err = c.portSetup()
	if err != nil {
		return c, err
	}
	c.pprofEnable, err = c.pprofEnableSetup()
	if err != nil {
		return c, err
	}
	return c, nil
}

func validateEnvVars(envVars map[string]string, requiredKeys []string) error {
	for _, key := range requiredKeys {
		if strings.TrimSpace(key) != "" {
			if _, exists := envVars[key]; !exists {
				return fmt.Errorf("missing mandatory env var: %q", key)
			}
		}
	}
	return nil
}
