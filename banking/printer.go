package banking

import "fmt"

type printer interface {
	println(text string)
}

type consolePrinter struct{}

func (c *consolePrinter) println(text string) {
	fmt.Println(text)
}

type testPrinter struct {
	text string
}

func (c *testPrinter) println(text string) {
	c.text = c.text + "\n" + text
}
