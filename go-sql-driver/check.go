package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	my_sql_password := os.Getenv("my_sql_password")
	if my_sql_password == "" {
		fmt.Println("my_sql_password enviroment variable is not, please set my_sql_password environment variable")
		return
	}
	db, err := sql.Open("mysql", "root:"+my_sql_password+"@tcp(localhost:3306)/testdb?charset=utf8")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM employeetest")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var first_name string
		var last_name string
		var age int
		var sex string
		var income float64
		err = rows.Scan(&first_name, &last_name, &age, &sex, &income)
		checkErr(err)
		fmt.Println(first_name, last_name, age, sex, income)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// go run check.go
// Mac Mohan 20 M 2000
// John Doe 20 M 50000
// John Doe 20 M 50000
// John Doe 20 M 50000
