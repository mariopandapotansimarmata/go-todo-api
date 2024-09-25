package repository

import (
	"context"
	"database/sql"
	"http-basic/helper"
	"http-basic/model/domain"
	"time"
)

type TodoImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoImpl{}
}

func (repository *TodoImpl) Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	query := "INSERT INTO todo (name, time_create) VALUES ($1, $2) RETURNING id"

	var id int64
	err := tx.QueryRowContext(ctx, query, todo.Name, todo.TimeCreate).Scan(&id)
	helper.PanicIfErr(err)

	todo.Id = int(id)

	return todo
}

func (repository *TodoImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	query := "UPDATE todo SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, todo.Name, todo.Id)
	helper.PanicIfErr(err)

	return todo
}

func (repository *TodoImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) {
	query := "DELETE from todo WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, todo.Id)
	helper.PanicIfErr(err)
}

func (repository *TodoImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	query := "SELECT * from todo WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, todoId)
	helper.PanicIfErr(err)
	defer rows.Close()

	todo := domain.Todo{}
	if rows.Next() {
		rows.Scan(&todo.Id, &todo.Name, &todo.TimeCreate, &todo.TimeFinish)
		return todo, nil
	} else {
		return todo, nil
	}
}

func (repository *TodoImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Todo, error) {
	query := "SELECT * from todo"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfErr(err)
	defer rows.Close()

	listTodos := []domain.Todo{}

	for rows.Next() {
		todo := domain.Todo{}
		rows.Scan(&todo.Id, &todo.Name, &todo.TimeCreate, &todo.TimeFinish)
		listTodos = append(listTodos, todo)
	}

	return listTodos, nil
}

func (repository *TodoImpl) SetFinish(ctx context.Context, tx *sql.Tx, todo domain.Todo, timeFinish time.Time) {
	query := "UPDATE todo SET time_finish = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, timeFinish, todo.Id)
	helper.PanicIfErr(err)
}
func (repository *TodoImpl) GetUsers(ctx context.Context, tx *sql.Tx, todo domain.Todo, timeFinish time.Time) {

}
