package service

import (
	"context"
	"database/sql"
	"fmt"
	"http-basic/helper"
	"http-basic/model/web"
	"http-basic/repository"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type AuthServiceImpl struct {
	AuthRepo repository.AuthRepo
	DB       *sql.DB
	Validate *validator.Validate
}

func NewAuthService(authRepo repository.AuthRepo, dB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{AuthRepo: authRepo, DB: dB, Validate: validate}
}

func (auth *AuthServiceImpl) GetUser(ctx context.Context, request web.AuthRequest) (web.AuthResponse, error) {
	err := auth.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := auth.DB.Begin()
	helper.PanicIfErr(err)

	defer helper.CommitOrRollBack(tx)

	user, err := auth.AuthRepo.GetUsers(ctx, tx, request.UserName)
	helper.PanicIfErr(err)

	if user.Username != request.UserName || user.Password != request.Password {
		return web.AuthResponse{}, fmt.Errorf("invalid username or password")
	}

	secretKey := []byte("this is very secret")
	expiresTime := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expiresTime,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return web.AuthResponse{}, err
	}

	return web.AuthResponse{
		Token:   tokenString,
		Expires: strconv.FormatInt(expiresTime, 10),
	}, nil
}
