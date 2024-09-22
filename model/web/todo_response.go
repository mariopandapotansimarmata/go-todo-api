package web

import "time"

type TodoResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	TimeCreate time.Time `json:"timeCreate"`
	TimeFinish time.Time `json:"timeFinish"`
}
