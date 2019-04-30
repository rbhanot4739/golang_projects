package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)
	go func() {
		for {
			c1 <- "Recieved from 1st Goroutine"
			time.Sleep(800 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Recieved from 2nd Goroutine"
			time.Sleep(2 * time.Second)
		}
	}()
	for {

		//fmt.Println(<-c1)
		//fmt.Println(<-c2)
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		default:
			fmt.Println("No msg recieved")
			time.Sleep(600 * time.Millisecond)

		}
	}
}
