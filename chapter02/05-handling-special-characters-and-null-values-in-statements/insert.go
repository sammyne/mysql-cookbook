package main

import (
	"fmt"

	"github.com/sammyne/mysql-cookbook/recipes/lib/cookbook"
)

func main() {
	db, err := cookbook.Connect()
	if nil != err {
		panic(err)
	}
	defer func() {
		if err := db.Close(); nil != err {
			panic(err)
		}
		fmt.Println("Disconnected")
	}()

	fmt.Println("Connected")

	stmt, err := db.Prepare(
		"INSERT INTO profile (name,birth,color,foods,cats) VALUES(?,?,?,?,?)")
	if nil != err {
		panic(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec("De'Mont", "1973-01-12", nil,
		"eggroll", 4); nil != err {
		panic(err)
	}
}
