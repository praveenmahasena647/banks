package db

import (
	"fmt"
)

func createAccount() {
	var _, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS account(
			id BIGSERIAL PRIMARY KEY NOT NULL,
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			created TIMESTAMP NOT NULL,
			email_ID VARCHAR(50) NOT NULL UNIQUE,
			account_type VARCHAR(70) NOT NULL,
			password VARCHAR(100) NOT NULL,
			email_varified BOOLEAN NOT NULL DEFAULT FALSE
		)
	`)
	fmt.Println(err)
}

func createCreditors() {
	var _, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS creditors(
			id BIGSERIAL PRIMARY KEY  NOT NULL,
			account_id BIGINT REFERENCES account(id),
			amount MONEY NOT NULL
		)
	`)
	fmt.Println(err)
}

func createDebitors() {
	var _, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS debitors(
				id BIGSERIAL PRIMARY KEY NOT NULL,
				account_id BIGINT REFERENCES account(id),
				amount MONEY NOT NULL
			);
		`)
	fmt.Println(err)
}
