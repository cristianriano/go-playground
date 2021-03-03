package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) introduce() {
	fmt.Println("My name is", p.name, "and I'm", p.age, "years old")
}

type people interface {
	introduce()
}

func structExample() {
	p := person{name: "Juan", age: 23}
	fmt.Println(p)

	p.age = 25
	fmt.Println(p.age)
}

func interfaceExample() {
	var p people
	p = person{name: "Cristian", age: 27}
	p.introduce()
}

// Use empty interface to hande any kind of value like Println
func emptyInterfaceExample() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func typeAssertionExample() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // 0, false
	fmt.Println(f, ok)

	// f = i.(float64) // panic a.k.a run time error
	// fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}