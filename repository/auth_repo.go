package repository

import (
	"context"
	"database/sql"
	"http-basic/model/domain"
)

type AuthRepo interface {
	GetUsers(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
