package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LoanController interface {
	Create(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	FindByKtp(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
}
