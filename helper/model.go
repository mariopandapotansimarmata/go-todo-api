package helper

import (
	"http-basic/model/domain"
	"http-basic/model/web"
)

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	return web.TodoResponse{
		Id:         todo.Id,
		Name:       todo.Name,
		TimeCreate: todo.TimeCreate,
		TimeFinish: todo.TimeFinish,
	}
}
