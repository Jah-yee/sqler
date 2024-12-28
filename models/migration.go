package models

import "time"

type MigrationEntity struct {
	ID         int64      `json:"id" db:"id"`
	Filename   string     `json:"filename" db:"filename"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	MigratedAt *time.Time `json:"migrated_at" db:"migrated_at"`
}
