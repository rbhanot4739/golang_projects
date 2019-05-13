package main

import (
	"fmt"
)

func main() {
	arr := []string{"Python", "Golang", "Perl", "Rust", "Erlang"}
	fmt.Println("Array is ", arr)
	if contains(arr, "Perl") {
		fmt.Println("String found in ", arr)
	}
}

func contains(a []string, str string) bool {
	for _, val := range a {
		if str == val {
			return true
		}
	}
	return false
}
