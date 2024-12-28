package contracts

import (
	"context"
	"github.com/alash3al/sqler/models"
)

type UserService interface {
	Create(context.Context, models.UserCreateInput) error
}
