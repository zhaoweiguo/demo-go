// usage: http://go-database-sql.org/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "<user>:<password>@tcp(<host>:<port>)/<db>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM <table>")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			//			fmt.Println("==", col)   // 都是byte[]
			if col == nil {
				value = "NUL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
		}
		fmt.Println("----------------")
	}

}
