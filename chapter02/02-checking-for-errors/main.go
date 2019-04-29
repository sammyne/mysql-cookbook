package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	const username = "baduser"
	const password = "badpass"
	const database = "cookbook"

	const dataSourceName = username + ":" + password + "@/" + database

	db, err := sql.Open("mysql", dataSourceName)
	if nil != err {
		if e, ok := err.(*mysql.MySQLError); ok {
			panic(fmt.Sprintf(
				"MySQL Error with error number as %d and description as %s",
				e.Number, e.Message))
		} else {
			panic(err)
		}
	}
	defer func() {
		db.Close()
		fmt.Println("Disconnected")
	}()

	fmt.Println("Connected")

	if err := db.Ping(); nil != err {
		if e, ok := err.(*mysql.MySQLError); ok {
			panic(fmt.Sprintf(
				"MySQL Error during ping with error number as %d and description as %s",
				e.Number, e.Message))
		} else {
			panic(err)
		}
	}
}
