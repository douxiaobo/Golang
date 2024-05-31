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
	rows, err := db.Query("SELECT * FROM test")
	checkErr(err)
	for rows.Next() {
		var id int
		var name string
		var age int
		var created_at string
		err = rows.Scan(&id, &name, &age, &created_at)
		checkErr(err)
		fmt.Println(id, name, age, created_at)
		// fmt.Println(id)
		// fmt.Println(name)
		// fmt.Println(age)
		// fmt.Println(created_at)
		defer rows.Close()

	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// douxiaobo@192 mattn_go-sqlite3 % go run query.go
// 1 Alice 25 2024-05-31T11:52:59Z
// 2 Tom 30 2024-05-31T12:02:32Z
// douxiaobo@192 mattn_go-sqlite3 % go run query.go
// 1
// Alice
// 25
// 2024-05-31T11:52:59Z
// 2
// Tom
// 30
// 2024-05-31T12:02:32Z
