package contracts

import (
	"context"
	"time"
)

type Migration interface {
	Prepare(ctx context.Context) error
	Status(ctx context.Context) ([]MigrationStatusOutputItem, error)
	Apply(ctx context.Context) ([]MigrationStatusOutputItem, error)
}

type MigrationStatusOutputItem struct {
	ID         int64      `json:"id" db:"id"`
	Filename   string     `json:"filename" db:"filename"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	MigratedAt *time.Time `json:"migrated_at" db:"migrated_at"`
}
