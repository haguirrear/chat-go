package service

import (
	"time"
)

// CreateUser

type CreateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateUserRes struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// LoginUser

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserRes struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
