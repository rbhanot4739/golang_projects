package main

// ArrayMultiplier is
func ArrayMultiplier(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 5
	}
	return arr
}
