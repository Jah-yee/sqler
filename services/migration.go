package services

import (
	"context"
	"github.com/alash3al/sqler/models"
)

type MigrationService interface {
	Prepare(ctx context.Context) error
	Status(ctx context.Context) ([]models.MigrationEntity, error)
	Apply(ctx context.Context) ([]models.MigrationEntity, error)
}
