package main

import (
	"fmt"
	"math"
)

func greet(name string) string {

	return "Hello " + name
}

func squares(x float64, y float64) (float64, float64) {
	return math.Pow(x, 2), math.Pow(y, 2)
}

func main() {

	fmt.Println(greet("Rohit"))
	num1, num2 := 3.0, 5.0
	fmt.Println(squares(num1, num2))
	nums := []int{2, 8, 12, 5}
	nums2 := ArrayMultiplier(nums)
	fmt.Println(nums2)
	i, j := 10, 30
	fmt.Printf("Original values of i, j ==> %d, %d\n", i, j)
	PointerPassing(&i, &j)
	fmt.Printf("Modified values of i, j after the fxn call==> %d, %d\n", i, j)

}
