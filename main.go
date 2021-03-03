package main

import (
	"fmt"
	"runtime"

	"github.com/cristianriano/go-playground/words"
)

func main() {
	fmt.Println("Running on: ", runtime.GOOS)
	fmt.Println(words.Count("My example"))
	// workerPoolWithCache()
}
