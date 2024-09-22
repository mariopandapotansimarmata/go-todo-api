package service

import (
	"context"
	"database/sql"
	"fmt"
	"http-basic/helper"
	"http-basic/model/domain"
	"http-basic/model/web"
	"http-basic/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	TodoRepo repository.TodoRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(todoRepo repository.TodoRepository, dB *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepo: todoRepo,
		DB:       dB,
		Validate: validate,
	}
}

func (todoService *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := todoService.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := todoService.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	todo := domain.Todo{
		Name:       request.Name,
		TimeCreate: time.Now(),
	}

	result := todoService.TodoRepo.Create(ctx, tx, todo)

	return helper.ToTodoResponse(result)
}

func (todoService *TodoServiceImpl) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	err := todoService.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := todoService.DB.Begin()
	helper.PanicIfErr(err)

	todo, err := todoService.TodoRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	todo.Name = request.Name
	fmt.Println(todo.Id)
	fmt.Println(todo.Name)

	result := todoService.TodoRepo.Update(ctx, tx, todo)

	defer helper.CommitOrRollBack(tx)
	return helper.ToTodoResponse(result)
}

func (todoService *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	tx, err := todoService.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	todo, err := todoService.TodoRepo.FindById(ctx, tx, todoId)
	helper.PanicIfErr(err)

	todoService.TodoRepo.Delete(ctx, tx, todo)
}

func (todoService *TodoServiceImpl) FindById(ctx context.Context, todoId int) web.TodoResponse {
	tx, err := todoService.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	todo, err := todoService.TodoRepo.FindById(ctx, tx, todoId)
	helper.PanicIfErr(err)

	return helper.ToTodoResponse(todo)
}

func (todoService *TodoServiceImpl) FindAll(ctx context.Context) []web.TodoResponse {
	tx, err := todoService.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	todos, err := todoService.TodoRepo.FindAll(ctx, tx)
	helper.PanicIfErr(err)

	listTodo := []web.TodoResponse{}

	for _, todo := range todos {
		listTodo = append(listTodo, helper.ToTodoResponse(todo))
	}

	return listTodo
}

func (todoService *TodoServiceImpl) SetFinish(ctx context.Context, request web.TodoSetFinishRequest, timeFinish time.Time) web.TodoResponse {
	err := todoService.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := todoService.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollBack(tx)

	todoResult, err := todoService.TodoRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	todoService.TodoRepo.SetFinish(ctx, tx, todoResult, timeFinish)

	todoResult, err = todoService.TodoRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	return helper.ToTodoResponse(todoResult)
}
