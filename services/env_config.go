package services

import (
	"fmt"
	"github.com/alash3al/sqler/contracts"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strings"
)

type EnvConfig struct {
	AppName              string `env:"APP_NAME,notEmpty,expand"`
	LogLevel             string `env:"LOG_LEVEL,notEmpty,expand"`
	HTTPServerListenHost string `env:"HTTP_SERVER_HOST,notEmpty,expand"`
	HTTPServerListenPort int    `env:"HTTP_SERVER_PORT,notEmpty,expand"`
	DatabaseDSN          string `env:"DATABASE_DSN,notEmpty,expand"`
}

func NewEnvConfig(envFilename string) (contracts.Config, error) {
	if _, err := os.Stat(envFilename); err == nil {
		if err := godotenv.Load(envFilename); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	var c EnvConfig

	if err := env.Parse(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c EnvConfig) GetAppName() string {
	return c.AppName
}

func (c EnvConfig) GetLogLevel() (slog.Leveler, error) {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return nil, fmt.Errorf("invalid log level specified: %s", c.LogLevel)
	}
}

func (c EnvConfig) GetHTTPServerHost() string {
	return c.HTTPServerListenHost
}

func (c EnvConfig) GetHTTPServerPort() int {
	return c.HTTPServerListenPort
}

func (c EnvConfig) GetDatabaseDriver() string {
	return "postgresql"
}

func (c EnvConfig) GetDatabaseDSN() string {
	return c.DatabaseDSN
}
