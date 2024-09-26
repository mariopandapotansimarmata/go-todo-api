package main

import (
	"fmt"
	"http-basic/controller"
	"http-basic/database"
	"http-basic/helper"
	"http-basic/middleware"
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
	authRepository := repository.NewAuthRepository()

	todoService := service.NewTodoService(todoRepository, db, validate)
	authService := service.NewAuthService(authRepository, db, validate)

	todoController := controller.NewTodoController(todoService)
	authController := controller.NewAuthController(authService)

	publicRouter := httprouter.New()
	publicRouter.POST("/api/v1/login", authController.SignIn)

	protectedRouter := httprouter.New()
	protectedRouter.GET("/api/v1/todos", todoController.FindAll)
	protectedRouter.POST("/api/v1/todos", todoController.Create)
	protectedRouter.GET("/api/v1/todos/:todoId", todoController.FindById)
	protectedRouter.PUT("/api/v1/todos/:todoId", todoController.Update)
	protectedRouter.DELETE("/api/v1/todos/:todoId", todoController.Delete)
	protectedRouter.PATCH("/api/v1/todos/:todoId/finish", todoController.SetFinish)

	publicHandler := middleware.EnableCORS(publicRouter)
	protectedHandler := middleware.EnableCORS(middleware.NewAuthMiddleware(protectedRouter)) // Apply authentication middleware here

	finalHandler := http.NewServeMux()

	finalHandler.Handle("/api/v1/login", publicHandler)
	finalHandler.Handle("/api/v1/", protectedHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: finalHandler,
	}

	fmt.Println("Web Server ready to serve...")
	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
