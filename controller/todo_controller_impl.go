package controller

import (
	"encoding/json"
	"http-basic/helper"
	"http-basic/model/web"
	"http-basic/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func (todoController *TodoControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	TodoCreateRequest := web.TodoCreateRequest{}
	err := decoder.Decode(&TodoCreateRequest)
	helper.PanicIfErr(err)

	todoResponse := todoController.TodoService.Create(r.Context(), TodoCreateRequest)

	webRepsonse := web.WebResponse{Code: http.StatusOK, Status: http.StatusText(http.StatusOK), Data: todoResponse}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webRepsonse)
	helper.PanicIfErr(err)

}

func (todoController *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (todoController *TodoControllerImpl) SetFinish(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
