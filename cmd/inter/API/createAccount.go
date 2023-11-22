package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/praveenmahasena647/bank/cmd/inter/types"
)

func createAccount(w http.ResponseWriter, r *http.Request) {
	var personData = types.Accounts{}
	var err = json.NewEncoder(w).Encode(&personData)
	fmt.Println(err, personData)
	fmt.Fprintln(w, "yayayayaaa")
}
