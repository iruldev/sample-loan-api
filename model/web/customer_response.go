package web

type CustomerResponse struct {
	Id   		int `json:"id"`
	Name 		string `json:"name"`
	Ktp			string `json:"ktp"`
	BirthDate	string `json:"birth_date"`
	Sex			int `json:"sex"`
}
