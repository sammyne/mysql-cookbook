// +build ignore

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

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

	var (
		id   int
		name string
		cats int
	)

	out := tabwriter.NewWriter(os.Stdout, 1, 2, 1, ' ', tabwriter.Debug)
	defer out.Flush()

	fmt.Fprintln(out, "id\tname\tcats")
	for rows.Next() {
		if err := rows.Scan(&id, &name, &cats); nil != err {
			panic(err)
		}

		//fmt.Println(id, name, cats)
		fmt.Fprintln(out, fmt.Sprintf("%d\t%s\t%d", id, name, cats))
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
}
