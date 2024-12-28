package contracts

import "log/slog"

type ConfigService interface {
	GetAppName() string
	GetLogLevel() (slog.Leveler, error)
	GetJWTSecret() string
	GetHTTPServerHost() string
	GetHTTPServerPort() int
	GetDatabaseDriver() string
	GetDatabaseDSN() string
}
