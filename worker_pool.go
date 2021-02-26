package main

import "fmt"

func workerPool() {
	jobs := make(chan int, 100)
	results := make(chan string, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i:= 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- string) {
	for n := range jobs {
		results <- fmt.Sprintf("%d => %d", n, fib(n))
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n - 1) + fib(n - 2)
}