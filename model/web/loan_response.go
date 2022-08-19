package web

type LoanResponse struct {
	Id   		int `json:"id"`
	CustomerId 	int `json:"customer_id"`
	Amount		int	`json:"amount"`
	Period		int `json:"period"`
	Purpose		string `json:"purpose"`
	Customer	CustomerResponse `json:"customer"`
}
