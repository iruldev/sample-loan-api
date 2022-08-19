package web

type CustomerCreateRequest struct {
	FirstName 	string	`validate:"required,max=225,min=1" json:"first_name"`
	LastName 	string	`validate:"required,max=225,min=1" json:"last_name"`
	Ktp			string	`validate:"required,alphanum,max=16,min=16" json:"ktp"`
	BirthDate	string	`validate:"required" json:"birth_date"`
	Sex			int		`validate:"required,number" json:"sex"`
}
