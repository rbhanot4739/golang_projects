package main

import "fmt"

func main() {
	// for loop
	for i := 1; i < 50; i++ {
		if i%3 == 0 {
			fmt.Printf("%d is divisble by 3\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d is divisble by 5\n", i)
		} else {
			fmt.Printf("%d is not divisble by 3 or 5\n", i)
		}
	}

	// while loop
	j := 0
	for j < 10 {
		fmt.Println(j * 3)
		j++
	}

	// infinite loop - loop without any condition
	// for {
	// 	fmt.Println("Hello")
	// }
	var color string
	fmt.Print("Enter a color: ")
	// reading input from user
	fmt.Scanln(&color)
	usingSwitch(color)
	usingSwitch("blue")
	usingSwitch("red")

	// for each loop

}

func usingSwitch(color string) {
	// syntax switch variable {cases option: ....}
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Blue color")
	default:
		fmt.Println("White is the default color")
	}
}
