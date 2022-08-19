package service

import (
	"context"
	"github.com/iruldev/sample-loan-api/model/web"
)

type LoanService interface {
	Create(ctx context.Context, request web.LoanCreateRequest) web.LoanResponse
	FindById(ctx context.Context, loanId int) web.LoanResponse
	FindByKtp(ctx context.Context, ktp string) web.CustomerLoanResponse
}
