package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Value = %v | Type = %T\n", i, i)
}

func checkType(i interface{}) {
	value, ok := i.(string)
	fmt.Println(value, ok)
}

func main() {
	i := 10
	describe(i)
	describe("hello")
	stu := struct {
		name string
		age  int
	}{"Johny", 77}
	describe(stu)

	checkType(78)
	checkType("Hello World")
}
