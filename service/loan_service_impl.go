package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/iruldev/sample-loan-api/exception"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/domain"
	"github.com/iruldev/sample-loan-api/model/web"
	"github.com/iruldev/sample-loan-api/repository"
)

type LoanServiceImpl struct {
	LoanRepository repository.LoanRepository
	CustomerService CustomerService
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewLoanService(
	loanRepository repository.LoanRepository,
	customerService CustomerService,
	DB *sql.DB,
	validate *validator.Validate) LoanService {
	return &LoanServiceImpl{
		LoanRepository: loanRepository,
		CustomerService: customerService,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *LoanServiceImpl) Create(ctx context.Context, request web.LoanCreateRequest) web.LoanResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := service.CustomerService.FirstOrSave(ctx, request.Customer)

	loan := domain.Loan{
		CustomerId: customer.Id,
		Amount:     int32(request.Amount),
		Period:     int8(request.Period),
		Purpose:    request.Purpose,
	}

	loan = service.LoanRepository.Save(ctx, tx, loan)

	return web.LoanResponse{
		Id:         loan.Id,
		CustomerId: loan.CustomerId,
		Amount:     int(loan.Amount),
		Period:     int(loan.Period),
		Purpose:    loan.Purpose,
		Customer:   web.CustomerResponse{
			Id:        customer.Id,
			Name:      customer.Name,
			Ktp:       customer.Ktp,
			BirthDate: customer.BirthDate,
			Sex:       customer.Sex,
		},
	}
}

func (service *LoanServiceImpl) FindById(ctx context.Context, loanId int) web.LoanResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	loan, err := service.LoanRepository.FindById(ctx, tx, loanId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer, err := service.CustomerService.FindById(ctx, loan.CustomerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.LoanResponse{
		Id:         loan.Id,
		CustomerId: loan.CustomerId,
		Amount:     int(loan.Amount),
		Period:     int(loan.Period),
		Purpose:    loan.Purpose,
		Customer:   web.CustomerResponse{
			Id:        customer.Id,
			Name:      customer.Name,
			Ktp:       customer.Ktp,
			BirthDate: customer.BirthDate,
			Sex:       customer.Sex,
		},
	}
}

func (service *LoanServiceImpl) FindByKtp(ctx context.Context, ktp string) web.CustomerLoanResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerService.FindByKtp(ctx, ktp)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	loans := service.LoanRepository.FindByCustomerId(ctx, tx, customer.Id)

	return web.CustomerLoanResponse{
		Id:        customer.Id,
		Name:      customer.Name,
		Ktp:       customer.Ktp,
		BirthDate: customer.BirthDate,
		Sex:       customer.Sex,
		Loans: loans,
	}
}
