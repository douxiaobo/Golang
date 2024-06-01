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
	stmt, err := db.Prepare("DELETE FROM mytable WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %d rows", affect)
}
