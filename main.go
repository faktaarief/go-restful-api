package main

import (
	"faktaarief/go-restful-api/app"
	"faktaarief/go-restful-api/controller"
	"faktaarief/go-restful-api/helper"
	"faktaarief/go-restful-api/middleware"
	"faktaarief/go-restful-api/repository"
	"faktaarief/go-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
