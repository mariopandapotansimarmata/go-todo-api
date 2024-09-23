package controller

import (
	"encoding/json"
	"http-basic/helper"
	"http-basic/model/web"
	"http-basic/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{TodoService: todoService}
}

func (todoController *TodoControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	TodoCreateRequest := web.TodoCreateRequest{}
	err := decoder.Decode(&TodoCreateRequest)
	helper.PanicIfErr(err)

	todoResponse := todoController.TodoService.Create(r.Context(), TodoCreateRequest)

	webRepsonse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todoResponse,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webRepsonse)
	helper.PanicIfErr(err)
}

func (todoController *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	todoUpdateRequest := web.TodoUpdateRequest{}
	err := decoder.Decode(&todoUpdateRequest)
	helper.PanicIfErr(err)

	todoId := params.ByName("todoId")

	id, err := strconv.Atoi(todoId)
	helper.PanicIfErr(err)

	todoUpdateRequest.Id = id

	todoResponse := todoController.TodoService.Update(r.Context(), todoUpdateRequest)

	webRepsonse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webRepsonse)
	helper.PanicIfErr(err)
}

func (todoController *TodoControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfErr(err)

	data := todoController.TodoService.FindById(r.Context(), id)
	todoController.TodoService.Delete(r.Context(), id)

	webRepsonse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webRepsonse)
	helper.PanicIfErr(err)
}

func (todoController *TodoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfErr(err)

	todoResponse := todoController.TodoService.FindById(r.Context(), id)

	webRepsonse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webRepsonse)
	helper.PanicIfErr(err)
}

func (todoController *TodoControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoResponses := todoController.TodoService.FindAll(r.Context())

	webRepsonses := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todoResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webRepsonses)
	helper.PanicIfErr(err)
}

func (todoController *TodoControllerImpl) SetFinish(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfErr(err)

	todoSetFinishRequest := web.TodoSetFinishRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&todoSetFinishRequest)
	todoSetFinishRequest.Id = id

	todoResponse := todoController.TodoService.SetFinish(r.Context(), todoSetFinishRequest, todoSetFinishRequest.TimeFinish)

	webRepsonse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todoResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err2 := encoder.Encode(webRepsonse)
	helper.PanicIfErr(err2)
}
