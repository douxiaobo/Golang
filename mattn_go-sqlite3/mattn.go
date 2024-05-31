package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// create table
	sqlStmt := `CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT, age INTEGER, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert data
	sqlStmt = `INSERT INTO test(name, age) VALUES(?,?)`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	name := "Alice"
	age := 25
	_, err = stmt.Exec(name, age)
	if err != nil {
		fmt.Println(err)
		return
	}

	// query data
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var created_at time.Time
		err = rows.Scan(&id, &name, &age, &created_at)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, name, age, created_at)
	}
}

// douxiaobo@192 mattn_go-sqlite3 % sqlite3 test.db
// SQLite version 3.43.2 2023-10-10 13:08:14
// Enter ".help" for usage hints.
// sqlite> .open test.db
// sqlite> .database
// main: /Users/douxiaobo/Documents/Coding/Practice in Coding/Golang/mattn_go-sqlite3/test.db r/w
// sqlite> select * from test;
// 1|Alice|25|2024-05-31 11:52:59
// sqlite> .quit
// douxiaobo@192 mattn_go-sqlite3 %
