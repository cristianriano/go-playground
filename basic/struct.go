package basic

import "fmt"

type person struct {
	name, lastname string
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

// EmptyInterfaceExample use empty interface to hande any kind of value like Println
func EmptyInterfaceExample() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// TypeAssertionExample attempts to cast
func TypeAssertionExample() {
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

// TypeSwitchExample using variable type on a switch
func TypeSwitchExample() {
	do(21)
	do("Hi")
	do(false)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}