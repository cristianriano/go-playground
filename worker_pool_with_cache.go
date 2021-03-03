package main

import (
	"fmt"
	"sync"
)

func workerPoolWithCache() {
	jobs := make(chan int, 10000)
	results := make(chan string, 10000)
	cache := map[int]int{0: 1, 1: 1}
	mutex := sync.RWMutex{}

	go workerWithCache(jobs, results, cache, &mutex)
	go workerWithCache(jobs, results, cache, &mutex)
	go workerWithCache(jobs, results, cache, &mutex)
	go workerWithCache(jobs, results, cache, &mutex)

	for i:= 0; i < 10000; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 10000; j++ {
		fmt.Println(<-results)
	}
}

func workerWithCache(jobs <-chan int, results chan<- string, cache map[int]int, mutex *sync.RWMutex) {
	for n := range jobs {
		res := fibWithCache(n, cache, mutex)

		mutex.Lock()
		cache[n] = res
		mutex.Unlock()

		results <- fmt.Sprintf("%d => %d", n, res)
	}
}

func fibWithCache(n int, cache map[int]int, mutex *sync.RWMutex) int {
	mutex.RLock()
	val := cache[n]
	mutex.RUnlock()

	if val != 0 {
		return val
	}


	return fibWithCache(n - 1, cache, mutex) + fibWithCache(n - 2, cache, mutex)
}