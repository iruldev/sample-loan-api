package app

import (
	"github.com/iruldev/sample-loan-api/controller"
	"github.com/iruldev/sample-loan-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(loanController controller.LoanController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/loans", loanController.Create)
	router.GET("/api/loans", loanController.FindByKtp)
	router.GET("/api/loans/:loanId", loanController.FindById)

	router.PanicHandler = exception.ErrorHandler
	return router
}
