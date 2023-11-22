package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DBConnect() error {
	var err error
	println("connecting to the DB")
	var psqlConfig = fmt.Sprintf("user=postgres password=CJ dbname=bank host=localhost port=5432 sslmode=disable")
	db, err = sql.Open("postgres", psqlConfig)
	if err != nil {
		return err
	}
	fmt.Println(err)
	return db.Ping()
}

func Seeding() error {
	createAccount()
	createCreditors()
	createDebitors()

	return nil
}
