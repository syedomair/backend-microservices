package envconstant

const (
	LogLevelEnvVar      = "LOG_LEVEL"
	DatabaseURLEnvVar   = "DATABASE_URL"
	PortEnvVar          = "PORT"
	DBEnvVar            = "DB"
	DBMaxIdleEnvVar     = "DB_MAX_IDLE"
	DBMaxOpenEnvVar     = "DB_MAX_OPEN"
	DBMaxLifeTimeEnvVar = "DB_MAX_LIFE_TIME"
	DBMaxIdleTimeEnvVar = "DB_MAX_IDLE_TIME"
	ZapConf             = "ZAP_CONF"
	GormConf            = "GORM_CONF"
	PprofEnable         = "PPROF_ENABLE"
)

const (
	Postgres = "POSTGRES"
	Mysql    = "MYSQL"
)
