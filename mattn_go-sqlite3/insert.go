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

	stmt, err := db.Prepare("INSERT INTO test(name,age) VALUES(?,?)")
	checkErr(err)
	res, err := stmt.Exec("Tom", 30)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("Last Insert Id:", id)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// douxiaobo@192 mattn_go-sqlite3 % go run insert.go
// Last Insert Id: 2
// douxiaobo@192 mattn_go-sqlite3 % sqlite3 test.db
// SQLite version 3.43.2 2023-10-10 13:08:14
// Enter ".help" for usage hints.
// sqlite> select * from test;
// 1|Alice|25|2024-05-31 11:52:59
// 2|Tom|30|2024-05-31 12:02:32
// sqlite> .quit
// douxiaobo@192 mattn_go-sqlite3 %
