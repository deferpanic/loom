package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "dbname=dptest sslmode=disable")

	var id int
	var sleep string

	err = db.QueryRow("select 1 as num, pg_sleep(2)").Scan(&id, &sleep)
	if err != nil {
		log.Println("oh no!")
	}

	/*
		err = db.QueryRow("select 1 as num, pg_sleep(3)").Scan(&id, &sleep)
		if err != nil {
			log.Println("oh no!")
		}
	*/
}
