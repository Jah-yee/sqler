package pgservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/alash3al/sqler/contracts"
	"github.com/alash3al/sqler/models"
	"github.com/alash3al/sqltmpl"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type JWTAuth struct {
	cfg  contracts.ConfigService
	pool *pgxpool.Pool
	tpl  *sqltmpl.Template
}

func NewJWTAuthService(c contracts.ConfigService, p *pgxpool.Pool, t *sqltmpl.Template) contracts.AuthService {
	return &JWTAuth{
		cfg:  c,
		pool: p,
		tpl:  t,
	}
}

func (j JWTAuth) AuthenticateViaEmailAndPassword(ctx context.Context, input models.AuthLoginWithEmailAndPasswordInput) (*models.AuthLoginOutput, error) {
	sql, args, err := j.tpl.Execute("user.findByEmail", input.Email)
	if err != nil {
		return nil, err
	}

	rows, err := j.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow[*models.UserEntity](rows, pgx.RowToAddrOfStructByName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return nil, nil
	}

	expiresAt := time.Now().Add(input.TTL)

	sql, args, err = j.tpl.Execute("auth.tokenCreate", models.AuthTokenCreateInput{
		OwnerType: "entity.user",
		OwnerID:   user.ID,
		ExpiresAt: expiresAt,
		UserAgent: input.UserAgent,
		IPAddress: input.IPAddress,
	})
	if err != nil {
		return nil, err
	}

	var tokenID int64

	if err := j.pool.QueryRow(ctx, sql, args...).Scan(&tokenID); err != nil {
		return nil, err
	}

	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprintf("%d", tokenID),
		Subject:   fmt.Sprintf("%d", user.ID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.cfg.GetJWTSecret()))
	if err != nil {
		return nil, err
	}

	return &models.AuthLoginOutput{
		User:  user,
		Token: token,
	}, nil
}

func (j JWTAuth) ValidateToken(ctx context.Context, token string) (*models.UserEntity, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.cfg.GetJWTSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userId, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, err
	}

	sql, args, err := j.tpl.Execute("user.findById", userId)
	if err != nil {
		return nil, err
	}

	rows, err := j.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgx.CollectOneRow[*models.UserEntity](rows, pgx.RowToAddrOfStructByName)
}
