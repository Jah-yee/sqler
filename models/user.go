package models

import "time"

type UserEntity struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	IsSysAdmin bool      `json:"is_sys_admin" db:"is_sys_admin"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type UserCreateInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsSysAdmin bool   `json:"is_sys_admin" db:"is_sys_admin"`
}
