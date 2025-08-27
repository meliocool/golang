package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

}

// * SQL Injection Solution -> Do not query SQL manually with String concatenation
// * Use parameter based SQL, and use functions like Execute or Query
func TestNoSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"

	// * SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)

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
}

// * INSERT INTO user(username, password) VALUES (?, ?)
func TestInsertSQLNoInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "gyj"
	password := "gyj11"
	sql_script := "INSERT INTO Customer(id, name) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, sql_script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert success!")
}

// * Auto Increment
// * There is a function result.LastInsertId() to get the latest inserted ID that was made with Auto Increment

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "deiton@gmail.com"
	comment := "Hello"

	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Comment with ID:", insertId)
}

// * Prepare Statement
// * We prep the query first, and then it is filled with parameter
// * Sometimes, we want to insert massive amount of data, like from a csv, with the same columns
// * The only difference would be the params, therefore we make a Prepare Statement

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// * db.Prepare(context, sql)
	// * Prepare statement is represented in a struct database/sql.Stmt
	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "dhitan" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment -" + strconv.Itoa(i)

		// Use the statement NOT the db
		res, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	// * Stmt must also be closed
	defer stmt.Close()
}

// * Database Transaction
// * By default every SQL query sent with GoLang is automatically Committed
// * If transaction is needed, use db.Begin() which results in a struct Tx (Transaction)
// * Tx will replace db, so Tx.Exec() or Tx.Query(), or tx.Prepare()
// * When done, use Tx.Commit() or Tx.Rollback()

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sqlQuery := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// * Do Transaction
	for i := 0; i < 10; i++ {
		email := "dhitan" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment -" + strconv.Itoa(i)

		res, err := tx.ExecContext(ctx, sqlQuery, email, comment)

		if err != nil {
			panic(err)
		}

		id, IDerr := res.LastInsertId()
		if IDerr != nil {
			panic(IDerr)
		}

		fmt.Println("Comment Id", id)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
}
