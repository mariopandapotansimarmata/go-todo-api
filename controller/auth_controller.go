package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	SignOut(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
