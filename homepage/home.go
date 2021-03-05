package homepage

import (
	"net/http"
)

// HomePageHandler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}