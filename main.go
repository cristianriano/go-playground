package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Running on: ", runtime.GOOS)
	fmt.Println(Count("My example"))
	// workerPoolWithCache()
}
