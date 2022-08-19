package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/domain"
)

type CustomerRepositoryImpl struct {

}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "insert into customers(name, ktp, birth_date, sex) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Ktp, customer.BirthDate, customer.Sex)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	customer.Id = int(id)

	return customer
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Customer, error) {
	SQL := "select id, name, ktp, birth_date, sex from customers where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}

	if rows.Next() {
		err = rows.Scan(&customer.Id, &customer.Name, &customer.Ktp, &customer.BirthDate, &customer.Sex)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("customer is not found")
	}
}


func (repository *CustomerRepositoryImpl) FindByKtp(ctx context.Context, tx *sql.Tx, ktp string) (domain.Customer, error) {
	SQL := "select id, name, ktp, birth_date, sex from customers where ktp = ?"
	rows, err := tx.QueryContext(ctx, SQL, ktp)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err = rows.Scan(&customer.Id, &customer.Name, &customer.Ktp, &customer.BirthDate, &customer.Sex)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("customer is not found")
	}
}
