package services

import (
	"context"
	"github.com/alash3al/sqler/models"
)

type AuthService interface {
	AuthenticateViaEmailAndPassword(ctx context.Context, input models.AuthLoginWithEmailAndPasswordInput) (*models.AuthLoginOutput, error)
	ValidateToken(ctx context.Context, token string) (*models.UserEntity, error)
}
