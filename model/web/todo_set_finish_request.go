package web

import "time"

type TodoSetFinishRequest struct {
	Id         int
	TimeFinish time.Time
}
