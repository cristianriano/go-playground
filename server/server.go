package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cristianriano/go-playground/homepage"
)

var (
	certFile 			 = os.Getenv("CERT_FILE")
	certKeyFile		 = os.Getenv("KEY_FILE")
)

const serviceAddress = ":8080"

func Start() {
	// Std flags AND file name/line
	logger := log.New(os.Stdout, "GO ", log.LstdFlags | log.Lshortfile)

	h := homepage.NewHandler(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := New(mux, serviceAddress)

	logger.Println("server starting")
	if certFile == "" || certKeyFile == "" {
		logger.Fatal(srv.ListenAndServe())
	} else {
		log.Fatal(srv.ListenAndServeTLS(certFile, certKeyFile))
	}
}

// New returns new server with tls config and timeouts
func New(mux *http.ServeMux, serverAddress string) *http.Server {
	// See https://blog.cloudflare.com/exposing-go-on-the-internet/ for details about these settings
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default cipher suite preferences, which are tuned to avoid attacks
		// Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	server := &http.Server{
		Addr:           serverAddress,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout: 		120 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      tlsConfig,
		Handler:        mux,
	}

	return server
}