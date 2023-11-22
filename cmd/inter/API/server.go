package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() error {
	var router = mux.NewRouter()

	router.HandleFunc("/createAccount", createAccount).Methods("POST")

	return http.ListenAndServe(":42069", corsHandler(router))
}
