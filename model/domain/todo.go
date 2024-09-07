package domain

import "time"

type Todo struct {
	Id         int
	Name       string
	TimeCreate time.Time
	TimeFinish time.Time
}
