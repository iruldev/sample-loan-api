package app

import (
	"database/sql"
	"github.com/iruldev/sample-loan-api/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		dataSourceName = "root:root@tcp(localhost:33061)/loan-api"
	}
	db, err := sql.Open("mysql", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
