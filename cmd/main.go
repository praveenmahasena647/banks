package main

import (
	"fmt"
	"os"

	db "github.com/praveenmahasena647/bank/cmd/inter/DB"
)

func main() {
	var DBConnectionErr = db.DBConnect()
	if DBConnectionErr != nil {
		fmt.Printf("%v", DBConnectionErr)
		os.Exit(1)
	}
	var seedlingErr error = db.Seedling()
	if seedlingErr != nil {
		fmt.Printf("%v", seedlingErr)
		os.Exit(1)
	}
}
