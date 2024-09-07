package web

import (
	"time"
)

type TodoSetFinishRequest struct {
	Id         int       `validate:"required"`
	TimeFinish time.Time `validate:"required"`
}
