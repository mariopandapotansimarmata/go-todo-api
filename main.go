package main

import (
	"http-basic/controller"
	"http-basic/database"
	"http-basic/helper"
	"http-basic/repository"
	"http-basic/service"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.NewDB()

	validate := validator.New()

	todoRepository := repository.NewTodoRepository()
	service := service.NewCategoryService(todoRepository, db, validate)
	controller := controller.NewTodoController(service)

	router := httprouter.New()

	router.GET("/api/v1/todos", controller.FindAll)
	router.POST("/api/v1/todos", controller.Create)
	router.GET(path("/todos/:todoId"), controller.FindById)
	router.PUT(path("/todos/:todoId"), controller.Update)
	router.DELETE(path("/todos/:todoId"), controller.Delete)
	router.PATCH(path("/todos/:todoId/"), controller.SetFinish)

	handler := enableCORS(router)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}

func path(addPath string) string {
	return "/api/v1" + addPath
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}
