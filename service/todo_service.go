package service

import (
	"context"
	"http-basic/model/web"
	"time"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) web.TodoResponse
	FindByAll(ctx context.Context) []web.TodoResponse
	SetFinish(ctx context.Context, request web.TodoSetFinishRequest, timeFinish time.Time) web.TodoResponse
}
