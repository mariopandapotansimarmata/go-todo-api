package repository

import (
	"context"
	"database/sql"
	"http-basic/model/domain"
)

type AuthRepoImpl struct {
}

func NewAuthRepository() AuthRepo {
	return &AuthRepoImpl{}
}

func (auth *AuthRepoImpl) GetUsers(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	query := "SELECT username, password FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		return domain.User{}, err
	}
	defer rows.Close()

	foundUser := domain.User{}

	if rows.Next() {
		if err := rows.Scan(&foundUser.Username, &foundUser.Password); err != nil {
			return domain.User{}, err
		}
		return foundUser, nil
	}

	return domain.User{}, sql.ErrNoRows
}
