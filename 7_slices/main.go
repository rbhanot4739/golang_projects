package main

import "fmt"

func main() {
	var arr1 [] int
	showArrDetails(arr1)
	arr1 = append(arr1, 2)
	showArrDetails(arr1)
	arr1 = append(arr1, 905, 3, 7, 6)
	showArrDetails(arr1)

	nums := make([]int, 5, 10)
	showArrDetails(nums)
}

func showArrDetails(arr []int) {
	fmt.Printf("Len = %d, Cap = %d, array = %d\n", len(arr), cap(arr), arr)
}
