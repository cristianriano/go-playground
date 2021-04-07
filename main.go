package main

import (
	"fmt"

	"github.com/cristianriano/go-playground/heap"
)

func main() {
	h := heap.MaxHeap{}

	h.Push(40)
	h.Push(30)
	h.Push(10)
	h.Push(15)
	h.Push(5)
	h.Push(20)
	h.Push(50)

	h.Print()

	fmt.Printf("%v, %v, %v", h.Pop(), h.Pop(), h.Pop())
	h.Print()
}