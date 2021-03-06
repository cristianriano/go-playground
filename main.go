package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cristianriano/go-playground/homepage"
	"github.com/cristianriano/go-playground/server"
)

var (
	certFile 			 = os.Getenv("CERT_FILE")
	certKeyFile		 = os.Getenv("KEY_FILE")
)

const serviceAddress = ":8080"

func main() {
	// Std flags AND file name/line
	logger := log.New(os.Stdout, "GO ", log.LstdFlags | log.Lshortfile)

	h := homepage.NewHandler(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, serviceAddress)

	logger.Println("server starting")
	if certFile == "" || certKeyFile == "" {
		logger.Fatal(srv.ListenAndServe())
	} else {
		log.Fatal(srv.ListenAndServeTLS(certFile, certKeyFile))
	}
}
