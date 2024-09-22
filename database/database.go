package database

import (
	"database/sql"
	"http-basic/helper"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewDB() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/todo"

	db, err := sql.Open("pgx", connStr)
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
