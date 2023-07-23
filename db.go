package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func connectToDB() (*sql.DB, error) {
	return sql.Open("mysql", os.Getenv("DB_URL"))
}
