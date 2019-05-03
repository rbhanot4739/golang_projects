package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Main")
	ss := []string{"one", "two"}
	var ms map[int]string
	ms[1] = "Hello"
	ms[2] = "World"
	t1 := reflect.TypeOf(ss).Kind()
	t2 := reflect.TypeOf(ms).Kind()
	if t1 == reflect.Slice {

	}
	if t2 == reflect.Map {

	}
}
