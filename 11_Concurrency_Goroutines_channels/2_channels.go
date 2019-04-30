package main

import (
	"fmt"
)

func main() {
	fmt.Println("Staring the main Goroutine")
	schan := make(chan int) // Unbuffered channel
	go sumSquares(234, schan)
	fmt.Println(<-schan) // reading from the channel
}

func sumSquares(number int, channel chan int) {
	var sum int
	for number != 0 {
		val := number % 10
		sum += val * val
		number = number / 10
	}
	channel <- sum // writing to the channel
}
