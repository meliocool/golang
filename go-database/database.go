package go_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// * Use Database

	db, err := sql.Open("mysql", "root:M3l1io0805@tcp(localhost:3306)/go-database?parseTime=true")
	if err != nil {
		panic(err)
	}

	// * There are a few functions to set database pooling

	// * db.SetMaxIdleConns(number) -> Minimum connection
	db.SetMaxIdleConns(10)

	// * db.SetMaxOpenConns(number) -> Maximum connection
	db.SetMaxOpenConns(100)

	// * db.SetConnMaxIdleTime(duration) -> How long an unused connection last before erased
	db.SetConnMaxIdleTime(5 * time.Minute)

	// * db.SetConnMaxLifetime(duration) -> How long a connection can be used
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
