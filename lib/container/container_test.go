package container

import (
	"maps"
	"os"
	"reflect"
	"testing"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Test_container_Db(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *gorm.DB
	}{
		{
			name: "Test Db function",
			fields: fields{
				db: &gorm.DB{},
			},
			want: &gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			if got := c.Db(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Db() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_Logger(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *zap.Logger
	}{
		{
			name: "Test Logger function",
			fields: fields{
				logger: &zap.Logger{},
			},
			want: &zap.Logger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			if got := c.Logger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_Port(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Port function",
			fields: fields{
				port: "8080",
			},
			want: "8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			if got := c.Port(); got != tt.want {
				t.Errorf("Port() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_PprofEnable(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test PprofEnable function",
			fields: fields{
				pprofEnable: "true",
			},
			want: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			if got := c.PprofEnable(); got != tt.want {
				t.Errorf("PprofEnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_getRequiredEnvVar(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test getRequiredEnvVar function with existing env var",
			fields: fields{
				environmentVariables: map[string]string{"TEST_ENV": "test_value"},
			},
			args: args{
				key: "TEST_ENV",
			},
			want:    "test_value",
			wantErr: false,
		},
		{
			name: "Test getRequiredEnvVar function with non-existing env var",
			fields: fields{
				environmentVariables: map[string]string{},
			},
			args: args{
				key: "NON_EXISTING_ENV",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			got, err := c.getRequiredEnvVar(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRequiredEnvVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRequiredEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		envVars map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    Container
		wantErr bool
	}{
		{
			name: "Test New function with all required env vars",
			args: args{
				envVars: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   os.Getenv(DatabaseURLEnvVar),
					GormConf:            "../../config/gorm-logger-config.json",
					ZapConf:             "../../config/zap-logger-config.json",
					PortEnvVar:          "8080",
					PprofEnable:         "true",
					DBEnvVar:            "postgres",
				},
			},
			want: &container{
				environmentVariables: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   os.Getenv(DatabaseURLEnvVar),
					GormConf:            "../../config/gorm-logger-config.json",
					ZapConf:             "../../config/zap-logger-config.json",
					PortEnvVar:          "8080",
					PprofEnable:         "true",
					DBEnvVar:            "postgres",
				},
			},
			wantErr: false,
		},
		{
			name: "Test New function with missing required env var",
			args: args{
				envVars: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   os.Getenv(DatabaseURLEnvVar),
					GormConf:            "../../config/gorm-logger-config.json",
					ZapConf:             "../../config/zap-logger-config.json",
					PortEnvVar:          "8080",
					DBEnvVar:            "postgres",
				},
			},
			want: &container{
				environmentVariables: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   os.Getenv(DatabaseURLEnvVar),
					GormConf:            "../../config/gorm-logger-config.json",
					ZapConf:             "../../config/zap-logger-config.json",
					PortEnvVar:          "8080",
					DBEnvVar:            "postgres",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.envVars)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != nil {
				c, ok := got.(*container)
				if !ok {
					t.Fatalf("New() returned unexpected type, want *container, got %T", got)
				}
				// check env vars are correctly set
				if !maps.Equal(c.environmentVariables, tt.want.(*container).environmentVariables) {
					t.Errorf("New() environmentVariables = %v, want %v", c.environmentVariables, tt.want.(*container).environmentVariables)
				}
				// check other fields are not nil, assuming they are correctly initialized in dbSetup, loggerSetup, etc.
				if c.db == nil {
					t.Error("New() db is nil")
				}
				if c.logger == nil {
					t.Error("New() logger is nil")
				}
				if c.port == "" {
					t.Error("New() port is empty")
				}
				if c.pprofEnable == "" {
					t.Error("New() pprofEnable is empty")
				}
			}
		})
	}
}

func Test_validateEnvVars(t *testing.T) {
	type args struct {
		envVars      map[string]string
		requiredKeys []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test validateEnvVars with all required keys present",
			args: args{
				envVars: map[string]string{
					"KEY1": "value1",
					"KEY2": "value2",
				},
				requiredKeys: []string{"KEY1", "KEY2"},
			},
			wantErr: false,
		},
		{
			name: "Test validateEnvVars with missing required key",
			args: args{
				envVars: map[string]string{
					"KEY1": "value1",
				},
				requiredKeys: []string{"KEY1", "KEY2"},
			},
			wantErr: true,
		},
		{
			name: "Test validateEnvVars with empty key in requiredKeys",
			args: args{
				envVars: map[string]string{
					"KEY1": "value1",
				},
				requiredKeys: []string{"KEY1", ""},
			},
			wantErr: false,
		},
		{
			name: "Test validateEnvVars with whitespace key in requiredKeys",
			args: args{
				envVars: map[string]string{
					"KEY1": "value1",
				},
				requiredKeys: []string{"KEY1", "   "},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateEnvVars(tt.args.envVars, tt.args.requiredKeys); (err != nil) != tt.wantErr {
				t.Errorf("validateEnvVars() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_container_dbSetup(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Successful DB Setup",
			fields: fields{
				environmentVariables: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   os.Getenv(DatabaseURLEnvVar),
					GormConf:            "../../config/zap-logger-config.json",
					DBEnvVar:            "postgres",
				},
				db: nil,
			},
			wantErr: false,
		},
		{
			name: "Missing DBMaxIdleEnvVar",
			fields: fields{
				environmentVariables: map[string]string{
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   "test_db_url",
					GormConf:            "../../config/zap-logger-config.json",
					DBEnvVar:            "postgres",
				},
				db: nil,
			},
			wantErr: true,
		},
		{
			name: "Invalid DBMaxIdleEnvVar",
			fields: fields{
				environmentVariables: map[string]string{
					DBMaxIdleEnvVar:     "invalid",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   "test_db_url",
					GormConf:            "../../config/zap-logger-config.json",
					DBEnvVar:            "postgres",
				},
				db: nil,
			},
			wantErr: true,
		},
		{
			name: "Missing DatabaseURLEnvVar",
			fields: fields{
				environmentVariables: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					GormConf:            "../../config/zap-logger-config.json",
					DBEnvVar:            "postgres",
				},
				db: nil,
			},
			wantErr: true,
		},
		{
			name: "Missing GormConf",
			fields: fields{
				environmentVariables: map[string]string{
					DBMaxIdleEnvVar:     "10",
					DBMaxOpenEnvVar:     "100",
					DBMaxLifeTimeEnvVar: "3600",
					DBMaxIdleTimeEnvVar: "600",
					DatabaseURLEnvVar:   "test_db_url",
					DBEnvVar:            "postgres",
				},
				db: nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			_, err := c.dbSetup()
			if (err != nil) != tt.wantErr {
				t.Errorf("dbSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_container_loggerSetup(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Successful Logger Setup",
			fields: fields{
				environmentVariables: map[string]string{
					ZapConf: "../../config/zap-logger-config.json",
				},
				logger: nil,
			},
			wantErr: false,
		},
		{
			name: "Missing ZapConf",
			fields: fields{
				environmentVariables: map[string]string{},
				logger:               nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			_, err := c.loggerSetup()
			if (err != nil) != tt.wantErr {
				t.Errorf("loggerSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_container_portSetup(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Successful Port Setup",
			fields: fields{
				environmentVariables: map[string]string{
					PortEnvVar: "8080",
				},
			},
			want:    "8080",
			wantErr: false,
		},
		{
			name: "Missing PortEnvVar",
			fields: fields{
				environmentVariables: map[string]string{},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			got, err := c.portSetup()
			if (err != nil) != tt.wantErr {
				t.Errorf("portSetup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("portSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_pprofEnableSetup(t *testing.T) {
	type fields struct {
		logger               *zap.Logger
		db                   *gorm.DB
		port                 string
		pprofEnable          string
		environmentVariables map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Successful PprofEnable Setup",
			fields: fields{
				environmentVariables: map[string]string{
					PprofEnable: "true",
				},
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "Missing PprofEnable",
			fields: fields{
				environmentVariables: map[string]string{},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &container{
				logger:               tt.fields.logger,
				db:                   tt.fields.db,
				port:                 tt.fields.port,
				pprofEnable:          tt.fields.pprofEnable,
				environmentVariables: tt.fields.environmentVariables,
			}
			got, err := c.pprofEnableSetup()
			if (err != nil) != tt.wantErr {
				t.Errorf("pprofEnableSetup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pprofEnableSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}
