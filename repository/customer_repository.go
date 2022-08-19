package repository

import (
	"context"
	"database/sql"
	"github.com/iruldev/sample-loan-api/model/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Customer, error)
	FindByKtp(ctx context.Context, tx *sql.Tx, ktp string) (domain.Customer, error)
}
