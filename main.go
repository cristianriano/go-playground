package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Running on: ", runtime.GOOS)
	// fmt.Println(tour.WordCount("My example"))

	// f := tour.Fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	// workerPoolWithCache()
	emptyInterfaceExample()
}
