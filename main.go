package main

import (
	"fmt"

	"github.com/cristianriano/go-playground/dailyinterview"
)

func main() {
	x := []int{9,9,9,9,9,9,9}
	fmt.Println(dailyinterview.IndicesOnSorted(x, 9))
}