package web

type LoanCreateRequest struct {
	Amount	int		`validate:"required,gte=1000000,lte=10000000" json:"amount"`
	Period	int		`validate:"required,lte=12" json:"period"`
	Purpose	string	`validate:"required,oneof=vacation renovation electronics wedding rent car investment" json:"purpose"`
	Customer CustomerCreateRequest
}
