package main

import (
	"github.com/todo_test/controller"
	"github.com/todo_test/db"
	"github.com/todo_test/repository"
	"github.com/todo_test/router"
	"github.com/todo_test/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
