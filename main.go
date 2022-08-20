package main

import (
	"fmt"
	"github.com/iruldev/sample-loan-api/app"
	"github.com/iruldev/sample-loan-api/controller"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/web"
	"github.com/iruldev/sample-loan-api/repository"
	"github.com/iruldev/sample-loan-api/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	validate.RegisterValidation("Alphanum", web.Alphanum)
	validate.RegisterValidation("IsFormatMatch", web.IsFormatMatch)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)

	loanRepository := repository.NewLoanRepository()
	loanService := service.NewLoanService(loanRepository, customerService, db, validate)
	loanController := controller.NewLoanController(loanService)

	router := app.NewRouter(loanController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
