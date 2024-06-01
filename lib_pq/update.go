package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=douxiaobo password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("UPDATE mytable SET name=$1 WHERE id=$2")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("John Doe Updated", 1)
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(affect)
	rows, err := db.Query("select * from mytable")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}
}

// douxiaobo@192 lib_pq % go run update.go
// 2024/06/01 18:50:07 pq: database "test" does not exist
// exit status 1
// dbname=test，改为dbname=testdb

// douxiaobo@192 lib_pq % go run update.go
// 2024/06/01 18:50:34 1
// 2024/06/01 18:50:34 sql: expected 2 arguments, got 1
// exit status 1
// stmt.Query改为db.Query

// douxiaobo@192 lib_pq % go run update.go
// 2024/06/01 18:52:36 1
// 2024/06/01 18:52:36 id: 2, name: Antiony Davis, age: 25
// 2024/06/01 18:52:36 id: 1, name: John Doe Updated, age: 30
