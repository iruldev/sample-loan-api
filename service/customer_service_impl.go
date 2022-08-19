package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/domain"
	"github.com/iruldev/sample-loan-api/model/web"
	"github.com/iruldev/sample-loan-api/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(
	customerRepository repository.CustomerRepository,
	DB *sql.DB,
	validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) Save(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := domain.Customer{
		Name:      request.FirstName + " " + request.LastName,
		Ktp:       request.Ktp,
		BirthDate: request.BirthDate,
		Sex:       int8(request.Sex),
	}

	customer = service.CustomerRepository.Save(ctx, tx, customer)

	return web.CustomerResponse{
		Id:        customer.Id,
		Name:      customer.Name,
		Ktp:       customer.Ktp,
		BirthDate: customer.BirthDate,
		Sex:       int(customer.Sex),
	}
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, id int) (web.CustomerResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return web.CustomerResponse{
		Id:        customer.Id,
		Name:      customer.Name,
		Ktp:       customer.Ktp,
		BirthDate: customer.BirthDate,
		Sex:       int(customer.Sex),
	}, nil
}

func (service *CustomerServiceImpl) FindByKtp(ctx context.Context, ktp string) (web.CustomerResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindByKtp(ctx, tx, ktp)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return web.CustomerResponse{
		Id:        customer.Id,
		Name:      customer.Name,
		Ktp:       customer.Ktp,
		BirthDate: customer.BirthDate,
		Sex:       int(customer.Sex),
	}, nil
}

func (service *CustomerServiceImpl) FirstOrSave(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	customer, err := service.FindByKtp(ctx, request.Ktp)
	if err != nil {
		customer = service.Save(ctx, request)
	}
	return customer
}
