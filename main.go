package main

import (
	"Data-Category/helper"
	"Data-Category/middleware"
	"net/http"

	_ "github.com/lib/pq"
)

func NewServer(am *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: am,
	}
}

func main() {
	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
