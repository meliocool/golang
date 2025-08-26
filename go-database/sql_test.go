package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// * Execute SQL (Non-Query) -> INSERT UPDATE DELETE
	// * Context is heavily used in Database
	ctx := context.Background()
	sql_script := "INSERT INTO Customer(id, name) VALUES('nisha', 'Nisha')"
	// * db.ExecContext(context, sql, params)
	_, err := db.ExecContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert success!")
}

// * ExecContext does not return a result/value, therefore not fit for READ statements such as SELECT
// * Instead use db.QueryContext(context, sql, params)

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}

	// * The result is a struct sql.Rows
	// * Rows  is used to iterate the result of a query

	// * use rows.Next() returns a boolean, to iterate through the rows
	for rows.Next() {
		var id, name string
		// * To read every data, use rows.Scan(columns...)
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}

	// * don't forget to rows.Close()
	defer rows.Close()
}

// * Column Data Type
// * DataType in SQL are represented as certain data types in GoLang
// * VARCHAR, CHAR -> string
// * INT, BIGINT -> int32, int64
// * FLOAT, DOUBLE -> float32, float64
// * BOOLEAN -> bool
// * DATE, DATETIME, TIME, TIMESTAMP -> time.Time

func TestComplexSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer"
	rows, err := db.QueryContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float32
		var birth_date, created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		fmt.Println("Birth Date: ", birth_date)
		fmt.Println("Created At: ", created_at)
		fmt.Println("Married: ", married)
	}
}

func TestSQLNullable(t *testing.T) {
	// * By default Driver MySQL for golang will query data type DATE, DATETIME, TIMESTAMP,, to []byte or []uint8
	// * Which can be converted to string and parsed to time.Time
	// * Add ?parseTime=true after the dbname in GetConnection

	// * Nullable Type
	// * GoLang does not understand the NULL Data Type
	// * Special for NULL, it will be troublesome if we use Scan()
	// * Use the special Nullable data types from the package database/sql

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM customer"
	rows, err := db.QueryContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("============")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date: ", birth_date.Time)
		}
		fmt.Println("Created At: ", created_at)
		fmt.Println("Married: ", married)
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//username := "admin"
	//password := "admin"

	// * SQL Injection is a Cyber Security threat
	// * The input would manipulate a SQL statement which could return data that they should not access
	usernameInject := "admin'; #"
	passwordInject := "wrong"
	sql_script := "SELECT username FROM user WHERE username = '" + usernameInject + "' AND password = '" + passwordInject + "' LIMIT 1 "
	fmt.Println(sql_script)
	rows, err := db.QueryContext(ctx, sql_script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login!: ", username)
	}

	defer rows.Close()

	// * SQL Injection Solution -> Do not query SQL manually with String concatenation
	// * Use parameter based SQL, and use functions like Execute or Query
}
