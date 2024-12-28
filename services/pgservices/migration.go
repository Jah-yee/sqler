package pgservices

import (
	"bytes"
	"context"
	"embed"
	"github.com/alash3al/sqler/contracts"
	"github.com/alash3al/sqler/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"io/fs"
	"path/filepath"
	"time"
)

type Migrations struct {
	assetsFS embed.FS
	pool     *pgxpool.Pool
}

func NewMigrationService(assetsFS embed.FS, pool *pgxpool.Pool) contracts.MigrationService {
	return &Migrations{
		assetsFS: assetsFS,
		pool:     pool,
	}
}

func (e Migrations) Prepare(ctx context.Context) error {
	migrationsTableCreation := `
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			filename TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP NOT NULL,
			migrated_at TIMESTAMP
		);
	`

	if _, err := e.pool.Exec(ctx, migrationsTableCreation); err != nil {
		return err
	}

	migrationFiles, err := fs.Glob(e.assetsFS, "assets/postgresql/migrations/*.sql")
	if err != nil {
		return err
	}

	for _, migrationFile := range migrationFiles {
		insertSQL := `INSERT INTO migrations (filename, created_at) VALUES ($1, CURRENT_TIMESTAMP) ON CONFLICT DO NOTHING`
		if _, err := e.pool.Exec(ctx, insertSQL, filepath.Base(migrationFile)); err != nil {
			return err
		}
	}

	return nil
}

func (e Migrations) Status(ctx context.Context) ([]models.MigrationEntity, error) {
	rows, err := e.pool.Query(ctx, `SELECT * FROM migrations ORDER BY filename`)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows[models.MigrationEntity](rows, pgx.RowToStructByName)
}

func (e Migrations) Apply(ctx context.Context) ([]models.MigrationEntity, error) {
	rows, err := e.pool.Query(ctx, `SELECT * FROM migrations WHERE migrated_at IS NULL ORDER BY filename`)
	if err != nil {
		return nil, err
	}

	queue, err := pgx.CollectRows[models.MigrationEntity](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	err = pgx.BeginFunc(ctx, e.pool, func(tx pgx.Tx) error {
		for i, item := range queue {
			contents, err := fs.ReadFile(e.assetsFS, filepath.Join("assets/postgresql/migrations", item.Filename))
			if err != nil {
				return err
			}

			for _, stmtBytes := range bytes.Split(contents, []byte(";")) {
				stmtBytes = bytes.TrimSpace(stmtBytes)

				if len(stmtBytes) == 0 {
					continue
				}

				if _, err := tx.Exec(ctx, string(stmtBytes)); err != nil {
					return err
				}
			}

			if _, err := tx.Exec(ctx, `UPDATE migrations SET migrated_at = CURRENT_TIMESTAMP WHERE id = $1`, item.ID); err != nil {
				return err
			}

			now := time.Now()
			queue[i].MigratedAt = &now
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return queue, nil
}
