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

	stmt, err := db.Prepare("UPDATE employeetest SET age =? WHERE sex =?")
	checkErr(err)
	res, err := stmt.Exec(20, "M")
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	defer stmt.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// my_sql_password enviroment variable is not, please set my_sql_password environment variable

// go run update.go
// 3
