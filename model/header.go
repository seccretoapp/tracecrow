package model

import "fmt"

// LogLevel representa o nível do log
type LogLevel string

const (
	LogLevelInfo  LogLevel = "INFO"
	LogLevelError LogLevel = "ERROR"
	LogLevelDebug LogLevel = "DEBUG"
)

// Environment representa o ambiente
type Environment string

const (
	EnvironmentProduction  Environment = "PRODUCTION"
	EnvironmentStaging     Environment = "STAGING"
	EnvironmentDevelopment Environment = "DEVELOPMENT"
)

// FromStringToLogLevel converte uma string para LogLevel
func FromStringToLogLevel(level string) (LogLevel, error) {
	switch level {
	case string(LogLevelInfo):
		return LogLevelInfo, nil
	case string(LogLevelError):
		return LogLevelError, nil
	case string(LogLevelDebug):
		return LogLevelDebug, nil
	default:
		return "", fmt.Errorf("invalid log level: %s", level)
	}
}

// FromStringToEnvironment converte uma string para Environment
func FromStringToEnvironment(env string) (Environment, error) {
	switch env {
	case string(EnvironmentProduction):
		return EnvironmentProduction, nil
	case string(EnvironmentStaging):
		return EnvironmentStaging, nil
	case string(EnvironmentDevelopment):
		return EnvironmentDevelopment, nil
	default:
		return "", fmt.Errorf("invalid environment: %s", env)
	}
}

type Header struct {
	ID            string
	Timestamp     int64
	HMAC          string
	LogLevel      LogLevel
	Environment   Environment
	CorrelationID string
	Signature     []byte
	Nonce         []byte
	PublicKey     []byte
}
