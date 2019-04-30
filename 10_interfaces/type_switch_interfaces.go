package main

import "fmt"

func typeChecker(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("This is a string data. Value => %v, Type => %T\n", i, i)
	case int:
		fmt.Printf("This is an int. Value => %v, Type => %T\n", i, i)
	default:
		fmt.Println("Can't determine type")

	}
}

func main() {
	typeChecker(300)
	typeChecker("Hello")
	typeChecker(30.78)

}
