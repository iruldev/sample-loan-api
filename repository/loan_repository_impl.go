package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/domain"
)

type LoanRepositoryImpl struct {
}

func NewLoanRepository() LoanRepository {
	return &LoanRepositoryImpl{}
}

func (repository *LoanRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, loan domain.Loan) domain.Loan {
	SQL := "insert into loans(customer_id, amount, period, purpose) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, loan.CustomerId, loan.Amount, loan.Period, loan.Purpose)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	loan.Id = int(id)
	return loan
}

func (repository *LoanRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, loanId int) (domain.Loan, error) {
	SQL := "select id, customer_id, amount, period, purpose from loans where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, loanId)
	helper.PanicIfError(err)
	defer rows.Close()

	loan := domain.Loan{}
	if rows.Next() {
		err := rows.Scan(&loan.Id, &loan.CustomerId, &loan.Amount, &loan.Period, &loan.Purpose)
		helper.PanicIfError(err)
		return loan, nil
	} else {
		return loan, errors.New("loan is not found")
	}
}

func (repository *LoanRepositoryImpl) FindByCustomerId(ctx context.Context, tx *sql.Tx, customerId int) []domain.Loan {
	SQL := "select id, customer_id, amount, period, purpose from loans where customer_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	var loans []domain.Loan

	for rows.Next() {
		loan := domain.Loan{}
		err := rows.Scan(&loan.Id, &loan.CustomerId, &loan.Amount, &loan.Period, &loan.Purpose)
		helper.PanicIfError(err)
		loans = append(loans, loan)
	}

	return loans
}
