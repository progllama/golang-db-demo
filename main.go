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

func main() {
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
}

func connects(db *sql.DB) bool {
	err := db.Ping()
	return err == nil
}
