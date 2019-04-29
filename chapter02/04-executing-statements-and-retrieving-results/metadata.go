// +build ignore

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

	rows, err := db.Query("SELECT id, name, cats FROM profile")
	if nil != err {
		panic(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if nil != err {
		panic(err)
	}
	fmt.Println(columns)

	if err = rows.Err(); err != nil {
		panic(err)
	}
}
