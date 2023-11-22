package main

import (
	"fmt"
	"os"

	api "github.com/praveenmahasena647/bank/cmd/inter/API"
)

func main() {
	/*	if DBConnectionErr := db.DBConnect(); DBConnectionErr != nil {
			fmt.Printf("error during connecting to the DB %v", DBConnectionErr)
			os.Exit(1)
		}

		if seedErr := db.Seeding(); seedErr != nil {
			fmt.Printf("error during creating db_table %v", seedErr)
			os.Exit(0)
		}
	*/

	if serverErr := api.RunServer(); serverErr != nil {
		fmt.Printf("there was a server spwan error %v", serverErr)
		os.Exit(0)
	}
}
