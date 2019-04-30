package main

import (
	"fmt"
)

// closures in Go

func outer() func(n int) {
	i := 1
	return func(n int) {
		for i <= 10 {
			fmt.Printf("%d x %d = %d\n", n, i, n*i)
			i++
		}
	}
}

func main() {

	// anonymous functions
	func(name string) {
		fmt.Println("Hello", name)
	}("Anonymous")

	mul9 := outer()
	mul9(9)

	mul23 := outer()
	mul23(23)
}
