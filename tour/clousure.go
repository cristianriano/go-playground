package tour

// Fibonacci returns a Clousure which calls a fibonacci
func Fibonacci() func() int {
	x, y := 0, 1

	return func() int {
		defer func() {
			x, y = y, x+y
		} ()
		return x
	}
}