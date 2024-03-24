package banking

import "fmt"

type printer interface {
	println(text string)
}

type consolePrinter struct{}

func (c *consolePrinter) println(text string) {
	fmt.Println(text)
}
