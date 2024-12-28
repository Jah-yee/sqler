package pgservices

import (
	"context"
	"github.com/alash3al/sqler/contracts"
	"github.com/alash3al/sqler/models"
	"github.com/alash3al/sqltmpl"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	pool *pgxpool.Pool
	tpl  *sqltmpl.Template
}

func NewUserService(p *pgxpool.Pool, t *sqltmpl.Template) contracts.UserService {
	return &User{
		pool: p,
		tpl:  t,
	}
}

func (u User) Create(ctx context.Context, input models.UserCreateInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hash)

	sql, args, err := u.tpl.Execute("user.create", input)
	if err != nil {
		return err
	}

	if _, err := u.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
