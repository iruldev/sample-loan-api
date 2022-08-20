package web

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/iruldev/sample-loan-api/helper"
	"regexp"
	"strconv"
	"time"
)

type CustomerCreateRequest struct {
	FirstName 	string	`validate:"required,max=225,min=1" json:"first_name"`
	LastName 	string	`validate:"required,max=225,min=1" json:"last_name"`
	Ktp			string	`validate:"required,Alphanum,IsFormatMatch,max=16,min=16" json:"ktp"`
	BirthDate	string	`validate:"required" json:"birth_date"`
	Sex			int		`validate:"required,number" json:"sex"`
}

func Alphanum(fl validator.FieldLevel) bool {
	fmt.Println()
	value := fl.Field().String()
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return true
}

func IsFormatMatch(fl validator.FieldLevel) bool {
	ktp := fl.Field().String()
	sex := fl.Top().FieldByIndex([]int{3,4})
	birthDate := fl.Top().FieldByIndex([]int{3,3})

	// Validation KTP Number
	parse, err := time.Parse("02-01-2006", birthDate.String())
	helper.PanicIfError(err)

	year := strconv.Itoa(parse.Year())[2:] // yyyy -> yy
	day := parse.Day()
	// if a woman day + 40
	if sex.Int() == 2 {
		day += 40
	}

	formatted := fmt.Sprintf("%02d%02d%s", day, parse.Month(), year)
	matched, _ := regexp.MatchString(formatted, ktp)
	if !matched {
		return false
	}
	return true
	// End Validation KTP Number
}
