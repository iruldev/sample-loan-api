package service

import (
	"context"
	"github.com/iruldev/sample-loan-api/model/web"
)

type CustomerService interface {
	Save(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	FindById(ctx context.Context, id int) (web.CustomerResponse, error)
	FindByKtp(ctx context.Context, ktp string) (web.CustomerResponse, error)
	FirstOrSave(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
}
