package models

import "time"

type AuthLoginWithEmailAndPasswordInput struct {
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	TTL       time.Duration `json:"ttl"`
	UserAgent string        `json:"user_agent"`
	IPAddress string        `json:"ip_address"`
}

type AuthLoginOutput struct {
	User  *UserEntity `json:"user"`
	Token string      `json:"token"`
}

type AuthTokenEntity struct {
	ID        int64     `json:"id" db:"id"`
	OwnerType string    `json:"owner_type" db:"owner_type"`
	OwnerID   string    `json:"owner_id" db:"owner_id"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	UserAgent string    `json:"user_agent" db:"user_agent"`
	IPAddress string    `json:"ip_address" db:"ip_address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AuthTokenCreateInput struct {
	OwnerType string    `json:"owner_type" db:"owner_type"`
	OwnerID   int64     `json:"owner_id" db:"owner_id"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	UserAgent string    `json:"user_agent" db:"user_agent"`
	IPAddress string    `json:"ip_address" db:"ip_address"`
}
