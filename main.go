package main

import (
	"log"
	"net/http"

	"github.com/cristianriano/go-playground/homepage"
	"github.com/cristianriano/go-playground/server"
)

const serviceAddress = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage.HomePageHandler)

	srv := server.New(mux, serviceAddress)
	log.Fatal(srv.ListenAndServe())
}
