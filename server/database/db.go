package database

import (
	"database/sql"

	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// GetConnection connecting to the database and returning connection...
func GetConnection() (*sql.DB, error) {

	dbDriver := "mysql"  // Driver name
	dbUser := "root"     // Username
	dbPass := "password" // Password
	dbName := "gopress"  // Database name

	// Connecting to database:
	// sql.Open("DRIVER", "username:password/@database)
	// db varible using with `database/sql` driver.
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True")
	return db, err
}
