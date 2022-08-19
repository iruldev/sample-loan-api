package integration_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/iruldev/sample-loan-api/app"
	"github.com/iruldev/sample-loan-api/controller"
	"github.com/iruldev/sample-loan-api/helper"
	"github.com/iruldev/sample-loan-api/model/domain"
	"github.com/iruldev/sample-loan-api/repository"
	"github.com/iruldev/sample-loan-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL_TEST"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

//var BaseURL = fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))
var BaseURL = fmt.Sprintf("http://localhost:3000")

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)

	loanRepository := repository.NewLoanRepository()
	loanService := service.NewLoanService(loanRepository, customerService, db, validate)
	loanController := controller.NewLoanController(loanService)
	router := app.NewRouter(loanController)
	return router
}

func truncateLoans() {
	db := setupTestDB()
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE loans")
	db.Exec("TRUNCATE customers")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func TestCreateLoanSuccess(t *testing.T) {
	truncateLoans()
	router := setupRouter()

	requestBody := strings.NewReader(`{
		"amount": 1000000,
		"period": 6,
		"purpose": "investment",
		"customer": {
			"first_name": "khoirul",
			"last_name": "setyo",
			"ktp": "1234560509947890",
			"birth_date": "05-09-1994",
			"sex": 1
		}
	}`)

	request := httptest.NewRequest(http.MethodPost, BaseURL + "/api/loans", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	assert.Equal(t, 1000000, int(responseBody["data"].(map[string]interface{})["amount"].(float64)))
	assert.Equal(t, 6, int(responseBody["data"].(map[string]interface{})["period"].(float64)))
	assert.Equal(t, "investment", responseBody["data"].(map[string]interface{})["purpose"])
}

func TestCreateLoanFailed(t *testing.T) {
	truncateLoans()
	router := setupRouter()

	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPost, BaseURL + "/api/loans", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestGetLoanByIdSuccess(t *testing.T) {
	truncateLoans()

	// Create Loan and Customer
	tx, _ := setupTestDB().Begin()
	customerRepository := repository.NewCustomerRepository()
	customer := customerRepository.Save(context.Background(), tx, domain.Customer{
		Name:      "Khoirul Setyo",
		Ktp:       "1234560509947890",
		BirthDate: "05-09-1994",
		Sex:       1,
	})

	loanService := repository.NewLoanRepository()
	loan := loanService.Save(context.Background(), tx, domain.Loan{
		CustomerId: customer.Id,
		Amount:     5000000,
		Period:     1,
		Purpose:    "investment",
	})
	tx.Commit()

	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/loans/"+strconv.Itoa(loan.Id), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, loan.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, customer.Id, int(responseBody["data"].(map[string]interface{})["customer_id"].(float64)))
}

func TestGetLoanByIdFailed(t *testing.T) {
	truncateLoans()
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/loans/404", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}

func TestListLoansByCustomerKtpSuccess(t *testing.T) {
	truncateLoans()

	// Create Loans and Customer
	tx, _ := setupTestDB().Begin()
	customerRepository := repository.NewCustomerRepository()
	customer := customerRepository.Save(context.Background(), tx, domain.Customer{
		Name:      "Khoirul Setyo",
		Ktp:       "1234560509947890",
		BirthDate: "05-09-1994",
		Sex:       1,
	})

	loanService := repository.NewLoanRepository()

	loan1 := loanService.Save(context.Background(), tx, domain.Loan{
		CustomerId: customer.Id,
		Amount:     5000000,
		Period:     1,
		Purpose:    "investment",
	})

	loan2 := loanService.Save(context.Background(), tx, domain.Loan{
		CustomerId: customer.Id,
		Amount:     6000000,
		Period:     6,
		Purpose:    "investment",
	})
	tx.Commit()

	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/loans?ktp=" + customer.Ktp, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	data := responseBody["data"].(map[string]interface{})
	loans := data["loans"].([]interface{})

	loanResponse1 := loans[0].(map[string]interface{})
	loanResponse2 := loans[1].(map[string]interface{})

	assert.Equal(t, loan1.Id, int(loanResponse1["id"].(float64)))
	assert.Equal(t, loan1.Amount, int32(loanResponse1["amount"].(float64)))

	assert.Equal(t, loan2.Id, int(loanResponse2["id"].(float64)))
	assert.Equal(t, loan2.Amount, int32(loanResponse2["amount"].(float64)))
}

func TestListLoansByCustomerKtpFailed(t *testing.T) {
	truncateLoans()
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/loans?ktp=404", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}