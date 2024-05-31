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
	query(db)
	stmt, err := db.Prepare("UPDATE test SET name=?,age=? WHERE id=1")
	checkErr(err)
	res, err := stmt.Exec("Alice", 25)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	query(db)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func query(db *sql.DB) {
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

// douxiaobo@192 mattn_go-sqlite3 % go run update.go
// 1 Alice 25 2024-05-31 11:52:59 +0000 UTC
// 2 Tom 30 2024-05-31 12:02:32 +0000 UTC
// 1
// 1 Alice1 26 2024-05-31 11:52:59 +0000 UTC
// 2 Tom 30 2024-05-31 12:02:32 +0000 UTC
