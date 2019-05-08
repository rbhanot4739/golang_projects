package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Main")
	// ss := []int{1, 2}
	ms := make(map[string]string)
	ms["1"] = "Hello"
	ms["2"] = "World"
	// if !checkSliceType(ss, reflect.String) {
	// 	fmt.Println("Fail")
	// }
	if !checkMapType(ms, reflect.Int, reflect.String) {
		fmt.Println("Fail")
	}

}

func checkSliceType(d interface{}, t reflect.Kind) bool {
	if reflect.TypeOf(d).Kind() == reflect.Slice {
		if reflect.TypeOf(d).Elem().Kind() == t {
			return true
		}
	}
	return false

}

func checkMapType(d interface{}, kt reflect.Kind, vt reflect.Kind) bool {
	if reflect.TypeOf(d).Kind() == reflect.Map {
		if reflect.TypeOf(d).Key().Kind() == kt && reflect.TypeOf(d).Elem().Kind() == vt {
			return true
		}
	}
	return false
}
