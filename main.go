package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/alash3al/sqler/handlers/console"
	"github.com/alash3al/sqler/services"
	"github.com/alash3al/sqler/services/pgservices"
	"github.com/alash3al/sqltmpl"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli/v3"
	"log/slog"
	"os"
	"text/template"
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

	pgSQLTmpl := sqltmpl.New(
		template.Must(
			template.New("sqler").
				Delims("$${{", "}}$$").
				ParseFS(assetsFs, "assets/postgresql/templates/*.tpl.sql"),
		),
		func(i int) string {
			return fmt.Sprintf("$%d", i)
		},
	)

	migrationService := pgservices.NewMigrationService(assetsFs, dbConnPool)
	authService := pgservices.NewJWTAuthService(cfg, dbConnPool, pgSQLTmpl)
	userService := pgservices.NewUserService(dbConnPool, pgSQLTmpl)

	if err := migrationService.Prepare(context.Background()); err != nil {
		logger.Error(err.Error())
		return
	}

	app := cli.Command{
		Name:    "sqler",
		Version: "v3.0.0",
		Commands: []*cli.Command{
			console.MigrationHandler(migrationService),
			console.MakeHandler(userService, authService),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err.Error())
	}
}
