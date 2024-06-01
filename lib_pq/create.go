package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=douxiaobo password=postgres dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("Error connecting to database: %v", err)
	}
	defer db.Close()
	sql := `DROP TABLE IF EXISTS mytable;
	CREATE TABLE mytable (
		id SERIAL NOT NULL,
		name TEXT NOT NULL,
		age INT
	);`
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatalln("Error getting rows affected: %v", err)
	}

	println("Table created successfully")
}

// douxiaobo@192 lib_pq % go run create.go
// Table created successfully
