package auth

import (
	"context"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type TokenInfo interface {
	GetAccessToken() string
	GetExpiresAt() int64
	EncodeToJson() ([]byte, error)
	GetTokenType() string
}

type Author interface {
	GenerateToken(ctx context.Context, userID string) (TokenInfo, error)
	DestroyToken(ctx context.Context, accessToken string) error
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	Release() error
}
