package main

import (
	"context"
	"embed"
	"github.com/alash3al/sqler/handlers/console"
	"github.com/alash3al/sqler/services"
	"github.com/alash3al/sqler/services/pgservices"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
)

//go:embed assets
var assetsFs embed.FS

func main() {
	cfg, err := services.NewEnvConfig(".env")
	if err != nil {
		panic(err.Error())
	}

	logLevel, err := cfg.GetLogLevel()
	if err != nil {
		panic(err.Error())
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevel,
		ReplaceAttr: nil,
	}))

	dbConnPool, err := pgxpool.New(context.Background(), cfg.GetDatabaseDSN())
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer dbConnPool.Close()

	migrationService := pgservices.NewMigrations(assetsFs, dbConnPool)

	if err := migrationService.Prepare(context.Background()); err != nil {
		logger.Error(err.Error())
		return
	}

	app := cli.Command{
		Name:    "sqler",
		Version: "v3.0.0",
		Commands: []*cli.Command{
			console.MigrationHandler(migrationService),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		logger.Error(err.Error())
	}
}
