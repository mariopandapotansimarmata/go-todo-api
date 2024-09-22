package web

import (
	"time"
)

type TodoSetFinishRequest struct {
	Id         int       `validate:"required" json:"id"`
	TimeFinish time.Time `validate:"required" json:"timeFinish"`
}
