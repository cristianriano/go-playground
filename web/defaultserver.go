package web

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const defaultPort = ":8080"

// DefaultServer starts a server from go http package and does 10 request concurrent
func DefaultServer() {
	go startServer()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			makeRequest()
			wg.Done()
		}()
	}
	wg.Wait()
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 1000)
		fmt.Fprintf(w, "Hello world, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(defaultPort, nil))
}

func makeRequest() {
	resp, err := http.Get(fmt.Sprintf("http://localhost%s", defaultPort))
	if err != nil {
		fmt.Println("Client error")
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
		return
	}

	fmt.Println(string(body))
}
