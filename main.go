package main

import (
	"coredemo/framework"
	"net/http"
)

func main() {
	Server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    "localhost:8080",
	}
	Server.ListenAndServe()
}
