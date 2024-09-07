package repository

import (
	"context"
	"database/sql"
	"errors"
	"http-basic/helper"
	"http-basic/model/domain"
	"time"
)

type TodoImpl struct {
}

func (repository *TodoImpl) Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	query := "INSERT INTO todo (name, time_create) values (?,?)"
	res, err := tx.ExecContext(ctx, query, todo.Name, todo.TimeCreate)
	helper.PanicIfErr(err)

	id, err := res.LastInsertId()
	helper.PanicIfErr(err)

	todo.Id = int(id)

	return todo
}

func (repository *TodoImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	query := "UPDATE todo SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, todo.Name, todo.Id)
	helper.PanicIfErr(err)

	return todo
}

func (repository *TodoImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) {
	query := "DELETE from todo WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, todo.Id)
	helper.PanicIfErr(err)
}

func (repository *TodoImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	query := "SELECT * from todo WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, todoId)
	helper.PanicIfErr(err)

	todo := domain.Todo{}
	if rows.Next() {
		rows.Scan(&todo.Id, &todo.Name, &todo.TimeCreate, &todo.TimeFinish)
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}
}

func (repository *TodoImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Todo, error) {
	query := "SELECT * todo"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfErr(err)

	listTodos := []domain.Todo{}

	if rows.Next() {
		for rows.Next() {
			todo := domain.Todo{}
			rows.Scan(&todo.Id, &todo.Name, &todo.TimeCreate, &todo.TimeFinish)
			listTodos = append(listTodos, todo)
		}
	} else {
		return []domain.Todo{}, errors.New("todo is empty")
	}

	return listTodos, nil
}

func (repository *TodoImpl) SetFinish(ctx context.Context, tx *sql.Tx, todo domain.Todo, timeFinish time.Time) {
	query := "UPDATE todo SET time_finish = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, timeFinish, todo.Id)
	helper.PanicIfErr(err)
}
