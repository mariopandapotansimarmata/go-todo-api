package web

import "time"

type TodoResponse struct {
	Id         int
	Name       string
	TimeCreate time.Time
	TimeFinish time.Time
}
