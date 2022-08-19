package repository

import (
	"context"
	"database/sql"
	"github.com/iruldev/sample-loan-api/model/domain"
)

type LoanRepository interface {
	Save(ctx context.Context, tx *sql.Tx, loan domain.Loan) domain.Loan
	FindById(ctx context.Context, tx *sql.Tx, loanId int) (domain.Loan, error)
	FindByCustomerId(ctx context.Context, tx *sql.Tx, customerId int) []domain.Loan
}
