package main

import "fmt"

type person struct {
	name string
	age int
}

func structExample() {
	p := person{name: "Juan", age: 23}
	fmt.Println(p)

	p.age = 25
	fmt.Println(p.age)
}