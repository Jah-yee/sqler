package contracts

import "log/slog"

type Config interface {
	GetAppName() string
	GetLogLevel() (slog.Leveler, error)
	GetHTTPServerHost() string
	GetHTTPServerPort() int
	GetDatabaseDriver() string
	GetDatabaseDSN() string
}
