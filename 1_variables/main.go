package main

import "fmt"

func main() {

	var text string = "Hello World" /* you don't need to mention type here */
	var num1 = 5
	const num2 = 300
	fmt.Println(text, num1, num2)
	text, num1 = "Hey !!", 100
	//num2 = 200 // can't reassign to a constant
	Joiner()
}
