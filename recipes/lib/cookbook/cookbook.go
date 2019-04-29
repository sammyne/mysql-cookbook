package cookbook

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // introduce the mysql driver
)

// Connect establishes a connection to the cookbook database, returning
// a connection object. Throw an exception if the connection
// cannot be established.
func Connect() (*sql.DB, error) {
	const username = "cbuser"
	const password = "cbpass"
	const database = "cookbook"

	const dataSourceName = username + ":" + password + "@/" + database

	db, err := sql.Open("mysql", dataSourceName)
	if nil != err {
		return nil, err
	}

	if err := db.Ping(); nil != err {
		return nil, err
	}

	return db, nil
}
