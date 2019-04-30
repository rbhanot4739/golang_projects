package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Main")
	go greet("Go")
	greet("Python")
	//go greet("Python")
	//time.Sleep(time.Second * 3)
}

func greet(msg string) {
	i := 1
	for {
		fmt.Println("Hello", msg)
		time.Sleep(time.Millisecond * 500)
		i++
	}
}
