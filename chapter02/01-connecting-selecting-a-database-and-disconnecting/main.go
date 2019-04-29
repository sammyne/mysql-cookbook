package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	const username = "cbuser"
	const password = "cbpass"
	const database = "cookbook"

	const dataSourceName = username + ":" + password + "@/" + database

	db, err := sql.Open("mysql", dataSourceName)
	if nil != err {
		panic(err)
	}
	defer func() {
		db.Close()
		fmt.Println("Disconnected")
	}()

	fmt.Println("Connected")
}
