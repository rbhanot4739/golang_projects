package main

import (
	"fmt"
	"time"

	//"time"
)

func main() {
	bufChan := make(chan bool)
	fmt.Println("Staring Main")
	go printer("Hello", bufChan)
	go printer("World", bufChan)
	<-bufChan
	fmt.Println("Runnning Main")

}

func printer(st string, c chan bool) {
	for i := 1; i < 10; i++ {
		fmt.Printf("%v %v\n", st, i)
		time.Sleep(500 * time.Millisecond)
	}
	c <- true
}
