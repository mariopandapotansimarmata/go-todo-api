package service

import (
	"context"
	"http-basic/model/web"
)

type AuthService interface {
	GetUser(ctx context.Context, request web.AuthRequest) (web.AuthResponse, error)
}
