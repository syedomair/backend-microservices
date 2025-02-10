package logger

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// Create a temporary file with a valid zap config
	validConfig := `{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`

	tmpFile, err := os.CreateTemp("", "zap-config-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(validConfig)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Test case 1: Valid config file
	t.Run("ValidConfig", func(t *testing.T) {
		logger, err := New(tmpFile.Name())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if logger == nil {
			t.Error("Expected a valid logger, got nil")
		}
	})

	// Test case 2: Invalid config file path
	t.Run("InvalidConfigFilePath", func(t *testing.T) {
		_, err := New("nonexistent.json")
		if err == nil {
			t.Error("Expected an error for invalid config file path, got nil")
		}
	})

	// Test case 3: Invalid JSON config
	invalidConfig := `{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}` // Missing closing brace

	tmpFileInvalid, err := os.CreateTemp("", "zap-config-invalid-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFileInvalid.Name())

	if _, err := tmpFileInvalid.Write([]byte(invalidConfig)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	t.Run("InvalidJSONConfig", func(t *testing.T) {
		_, err := New(tmpFileInvalid.Name())
		if err == nil {
			t.Error("Expected an error for invalid JSON config, got nil")
		}
	})
}
