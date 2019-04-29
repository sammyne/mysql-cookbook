// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.NewConfig()

	config.User = "cbuser"
	config.Passwd = "cbpass"
	config.DBName = "cookbook"
	config.ParseTime = true // for parsing DATE/DATETIME into go's time.Time

	db, err := sql.Open("mysql", config.FormatDSN())
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

	stmt, err := db.Prepare("SELECT * FROM profile WHERE cats > ?")
	if nil != err {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(2)
	if nil != err {
		panic(err)
	}
	defer rows.Close()

	var (
		id    int
		name  string
		birth time.Time
		color sql.NullString
		foods string
		cats  int
	)

	out := tabwriter.NewWriter(os.Stdout, 1, 2, 1, ' ', tabwriter.Debug)
	defer out.Flush()

	fmt.Fprintln(out, "id\tname\tbirth\tcolor\tfoods\tcats")
	for rows.Next() {
		if err := rows.Scan(&id, &name, &birth, &color, &foods, &cats); nil != err {
			panic(err)
		}

		var c string
		if color.Valid {
			c = color.String
		} else {
			c = "NULL"
		}

		fmt.Fprintln(out, fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%d",
			id, name, birth, c, foods, cats))
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
}
