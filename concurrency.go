package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrency() {
	// Wait Group
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		// Reduce the counter
		wg.Done()
	}()

	wg.Wait()

	// Channel
	c := make(chan string)
	go countWithChannel("fish", c)

	for msg := range c {
		// Manually check if the channel is close
		// msg, open := <- c

		// if !open {
		// 	break
		// }

		fmt.Println("With channel", msg)
	}

	// Channel capacity
	cc := make(chan string, 2)
	cc <- "one"
	cc <- "two"

	msg2 := <- cc
	fmt.Println(msg2)

	msg2 = <- cc
	fmt.Println(msg2)

	// Channel select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		i := 0
		for i < 2 {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
			i++
		}
	}()

	go func() {
		i := 0
		for i < 2 {
			c2 <- "Every second"
			time.Sleep(time.Second)
			i++
		}
	}()

	i := 0
	for i < 4 {
		select {
		case m1 := <- c1:
			fmt.Println(m1)
		case m2 := <- c2:
			fmt.Println(m2)
		default:
			fmt.Println("...")
			time.Sleep(time.Millisecond * 250)
		}
		i++
	}
}

func count(thing string) {
	for i:= 0; i < 2; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countWithChannel(thing string, c chan string) {
	for i:= 0; i < 2; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}