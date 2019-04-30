package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting main Goroutine")
	chan1 := make(chan int)
	go sqrSum(450, chan1)
	fmt.Println("Sum of Squares of 450 is =", <-chan1)
	chan2 := make(chan int)
	go cubeSum(342, chan2)
	fmt.Println("Sum of cubes of 342 is ", <-chan2)
	fmt.Println("Finishing Main....")
}

func cubeSum(number int, shan chan int) {
	fmt.Println("Entering Cubes GoRoutine")
	sschan := make(chan int)
	sum := 0
	go makeDigits(number, sschan) // calling another goroutine inside a goroutine
	for val := range sschan {
		sum += val * val * val
	}
	shan <- sum
	fmt.Println("Exiting Cubes GoRoutine")
}

func sqrSum(number int, schan chan int) {
	fmt.Println("Entering Squares GoRoutine")
	sschan := make(chan int)
	sum := 0
	go makeDigits(number, sschan)
	for v := range sschan {
		sum += v * v
	}
	time.Sleep(3 * time.Second)
	schan <- sum
	fmt.Println("Exiting Squares GoRoutine")
}

func makeDigits(number int, dchan chan int) {
	fmt.Println("Entering intermediate GoRoutine")
	for number != 0 {
		val := number % 10
		dchan <- val
		number /= 10
	}
	close(dchan)
	fmt.Println("Exiting intermediate GoRoutine")
}
