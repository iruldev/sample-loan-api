package web

import "github.com/iruldev/sample-loan-api/model/domain"

type CustomerLoanResponse struct {
	Id   		int `json:"id"`
	Name 		string `json:"name"`
	Ktp			string `json:"ktp"`
	BirthDate	string `json:"birth_date"`
	Sex			int `json:"sex"`
	Loans		[]domain.Loan `json:"loans"`
}
