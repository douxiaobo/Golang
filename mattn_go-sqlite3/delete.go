package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM test WHERE id = ?")
	checkErr(err)
	res, err := stmt.Exec(3)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect, "rows deleted")
	defer stmt.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
