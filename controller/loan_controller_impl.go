package controller

import (
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/web"
	"github.com/iruldev/sample-loan-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LoanControllerImpl struct {
	LoanService service.LoanService
}

func NewLoanController(loanService service.LoanService) LoanController {
	return &LoanControllerImpl{
		LoanService: loanService,
	}
}

func (controller *LoanControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loanCreateRequest := web.LoanCreateRequest{}
	helper.ReadFromRequestBody(request, &loanCreateRequest)

	loanResponse := controller.LoanService.Create(request.Context(), loanCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loanResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *LoanControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loanId := params.ByName("loanId")

	id, err := strconv.Atoi(loanId)
	helper.PanicIfError(err)

	loanResponse := controller.LoanService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loanResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *LoanControllerImpl) FindByKtp(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ktp := request.URL.Query().Get("ktp")

	loanResponse := controller.LoanService.FindByKtp(request.Context(), ktp)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loanResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}
