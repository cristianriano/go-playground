package basic

import (
	"fmt"
)

func pointers() {
	x := 5
	incr(&x)
	fmt.Println(x)
}

func incr(x *int) {
	*x++
}