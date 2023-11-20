package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "CJ"
	dbname   = "bank"
	DB       *sql.DB
)

func DBConnect() error {
	var err error
	println("connecting to the DB")
	var psqlConfig = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlConfig)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func Seedling() error {
	var wg = sync.WaitGroup{}
	var funcCount int = 1 // incr if you add more functions
	var errStr = make(chan string)
	for i := 0; i < funcCount; i++ {
		wg.Add(1)
		go createAccount(&wg, errStr)
	}
	wg.Wait()
	if err := <-errStr; err != "" {
		return fmt.Errorf(err)
	}
	return nil
}

func createAccount(wg *sync.WaitGroup, e chan<- string) {
	var _, err = DB.Exec(`
	CREATE TABLE IF NOT EXIST account(
		id BIGSERIAL NOT NULL PRIMARY KEY,
		first_name VARCHAR(50) NOT NULL,
		last_name VARCHAR(50) NOT NULL,
		age smallint NOT NULL,
		email_ID VARCHAR(50),
		address VARCHAR(50) NOT NULL,
		account_type CHAR(10) NOT NULL,
	)
	`)
	if err != nil {
		println(err.Error())
		e <- err.Error()
	}
	defer wg.Done()
}
