package main

import (
	"Go-crud-api/helper"
	log "Go-crud-api/log"
	"Go-crud-api/v0/controller"
	"Go-crud-api/v0/repository"
	"Go-crud-api/v0/service"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
)

func KonekDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_crud_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}

func main() {

	db := KonekDB()

	e := echo.New()

	e.Use(log.MiddlewareLogging)
	e.HTTPErrorHandler = log.ErrorHandler

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories:categoryId", categoryController.Delete)

	router.POST("/api/products", productController.CreateProduct)

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	}

	fmt.Println("server start at localhost:9000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
