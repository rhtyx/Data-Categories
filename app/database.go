package app

import (
	"Data-Category/helper"
	"database/sql"
	"fmt"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "developer_category"
	password = "developer_only"
	dbName   = "data"
)

func NewDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
