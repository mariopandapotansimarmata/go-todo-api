package controller

import (
	"encoding/json"
	"http-basic/helper"
	"http-basic/model/web"
	"http-basic/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{authService: authService}
}

func (a AuthControllerImpl) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	authRequest := web.AuthRequest{}
	err := decoder.Decode(&authRequest)
	helper.PanicIfErr(err)

	response, err := a.authService.GetUser(r.Context(), authRequest)
	helper.PanicIfErr(err)

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	helper.PanicIfErr(err)

}

func (a AuthControllerImpl) SignOut(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
