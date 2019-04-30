package main

import "fmt"

func main() {
	var nums [3] int64
	nums[1] = 32
	nums[2] = 90000000000
	nums[0] = 100

	fmt.Println(nums)

	nums2 := [3]int16{1, 2, 3}
	fmt.Println(nums2)
	countries := [...]string{"India", "Uk", "Japan", "US"}
	fmt.Println(countries)
	var nums4 [] int
	nums4 = [] int{1, 2, 3, 5, 6,}
	fmt.Println(nums4)
	nums3 := [] int64{200, 400, 500, 900, 100, 200}
	fmt.Println(nums3[1:4])
}
