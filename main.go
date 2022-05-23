package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
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
		log.Println("Succeed to connect.")
	} else {
		log.Fatal("Fail to connect.")
	}

	names, err := getDatabaseNames(db)
	if err != nil {
		log.Fatal("Fail to get database names.")
	}
	log.Println(names)

	tableName := "sample"
	exist := false
	for _, v := range names {
		exist = exist || v == tableName
	}

	if !exist {
		err = createDatabase(db, tableName)
		if err != nil {
			log.Fatal(err)
		}
	}

	db.Exec("USE sample")

	sql := `create table Sample (
id integer, 
name varchar(10)
);`

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func connects(db *sql.DB) bool {
	err := db.Ping()
	return err == nil
}

func getDatabaseNames(db *sql.DB) ([]string, error) {
	//ここで作ると成功したとしても毎回からスライスが作成されるので良くない。
	empty := make([]string, 0)

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return empty, err
	}

	databases, err := tx.Query("SELECT datname FROM pg_database;")
	if err != nil {
		tx.Rollback()
		return empty, err
	}

	names := make([]string, 0)
	for databases.Next() {
		var tableName string
		databases.Scan(&tableName)
		names = append(names, tableName)
	}

	err = tx.Commit()
	if err != nil {
		return empty, err
	}

	return names, err
}

func createDatabase(db *sql.DB, tableName string) error {
	_, err := db.Exec("CREATE DATABASE " + tableName)
	return err
}
