package go_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	// * Creating a connection to the Database
	// * In GoLang, to do this, create an object sql.DB using the function sql.Open(driver, dataSourceName)
	// * For MySQL, the dataSourceName looks like username:password@tcp(host:port)/database_name
	// * This creates a POOL, the pooling in GoLang is automatic, which could be used to set the minimum and maximum connection
	db, err := sql.Open("mysql", "root:M3l1io0805@tcp(localhost:3306)/go-database")
	if err != nil {
		panic(err)
	}

	// * Use the Close() function when done
	defer db.Close()
}
