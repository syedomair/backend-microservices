package container

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func NewLogger(zapConfig string) (*zap.Logger, error) {

	logger := &zap.Logger{}
	file, err := os.Open(zapConfig)
	if err != nil {
		return logger, fmt.Errorf("failed to open logger config file")
	}
	defer file.Close()

	var cfg zap.Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse logger config json")
	}

	if logger, err = cfg.Build(); err != nil {
		return nil, err
	}

	defer logger.Sync()

	logger.Debug("logger construction succeeded")
	return logger, nil
}
