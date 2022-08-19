package main

import (
	"github.com/iruldev/sample-loan-api/app"
	"github.com/iruldev/sample-loan-api/controller"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/repository"
	"github.com/iruldev/sample-loan-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)

	loanRepository := repository.NewLoanRepository()
	loanService := service.NewLoanService(loanRepository, customerService, db, validate)
	loanController := controller.NewLoanController(loanService)

	router := app.NewRouter(loanController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
