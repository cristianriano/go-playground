package main

import (
	"errors"
	"fmt"
	"math"
)

func hello() {
	// Variable declaration
	var x int
	x = 5
	y := 10

	// If statement
	if x + y < 5 {
		fmt.Println("Menor")
	} else if x == 5 {
		fmt.Println("Igual")
	} else {
		fmt.Println("Mayor")
	}

	// Array and Slices
	// var a [5]int
	a := []int{0,1,2,3,4}
	a[2] = 1
	a = append(a, 10)

	fmt.Println(a)

	// Map
	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	delete(vertices, "triangle")

	fmt.Println(vertices)

	// Iterators
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for index, value := range a {
		fmt.Println(index, value)
	}

	for key, value := range vertices {
		fmt.Println(key, value)
	}

	// Functions
	fmt.Println(sum(2,3))

	result, err := sqrt(-16)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func sum(x int, y int) int {
	return x + y;
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative")
	}

	return math.Sqrt(x), nil
}