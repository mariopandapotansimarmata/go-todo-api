package repository

import (
	"context"
	"database/sql"
	"http-basic/model/domain"
	"time"
)

type TodoRepository interface {
	Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Todo, error)
	SetFinish(ctx context.Context, tx *sql.Tx, todoId int, timeFinish time.Time)
}
