package homepage

import (
	"log"
	"net/http"
	"time"
)

// Handler struct defines the middlewares
type Handler struct {
	logger *log.Logger
}

// Home handles a http request
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}

// Logger middleware
func (h *Handler) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

// SetupRoutes mounts the handler in the route
func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

// NewHandler returns a handler with the logger dependency injected
func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
