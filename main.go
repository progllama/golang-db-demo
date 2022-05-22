package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "testuser"
	password = "testpass"
	dbname   = "test"
)

var (
	driver = "postgres"
	dsn    = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

var (
	id   int
	name string
)

func main() {
	log.SetFlags(log.Llongfile)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if connects(db) {
		fmt.Println("Succeed to connect")
	} else {
		log.Fatal("Fail to connect")
	}

	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func connects(db *sql.DB) bool {
	err := db.Ping()
	return err == nil
}
