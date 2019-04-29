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

}
