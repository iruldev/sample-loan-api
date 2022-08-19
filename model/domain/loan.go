package domain

type Loan struct {
	Id         	int `json:"id"`
	CustomerId	int `json:"customer_id"`
	Amount		int32 `json:"amount"`
	Period 		int8 `json:"period"`
	Purpose		string `json:"purpose"`
}
