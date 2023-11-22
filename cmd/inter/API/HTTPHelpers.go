package api

import (
	"github.com/gorilla/handlers"
)

var corsHandler = handlers.CORS(
	handlers.AllowedHeaders([]string{"Content-Type", "X-UserName", "X-UserEmail", "X-Jwt", "X-links", "X-Admin"}),
	handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"}),
	handlers.AllowedOrigins([]string{"*"}),
)
