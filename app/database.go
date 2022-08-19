package app

import (
	"database/sql"
	"github.com/iruldev/sample-loan-api/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
