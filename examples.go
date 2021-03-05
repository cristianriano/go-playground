package main

import (
	"fmt"
	"runtime"

	"github.com/cristianriano/go-playground/basic"
	"github.com/cristianriano/go-playground/concurrency"
	"github.com/cristianriano/go-playground/tour"
)

func otherExamples() {
	fmt.Println("Running on: ", runtime.GOOS)
	// fmt.Println(tour.WordCount("My example"))

	// f := tour.Fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	basic.WorkerPoolWithCache()
	tour.RunWebCrawler()
	concurrency.ProducerExample()
}
