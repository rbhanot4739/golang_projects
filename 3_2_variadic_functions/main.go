package main

import (
	"fmt"
)

//variadic functions are the fxns which can take multiple args

func multiplier(nums ...int) int {
	fmt.Printf("Type of nums is %T\n", nums)
	product := 1
	for _, val := range nums {
		product *= val
	}
	return product
}

func main() {
	fmt.Println(multiplier(1, 2, 3))
	slice1 := [] int{22, 33, 55}
	// the way to pass a slice to a variadic function is using '...' pack operator
	fmt.Println(multiplier(slice1...))
}
