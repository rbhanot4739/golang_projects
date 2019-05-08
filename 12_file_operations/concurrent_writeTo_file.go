package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Main GOroutine")
	var wg sync.WaitGroup

	// f, err := os.OpenFile("random_numbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	f, err := os.Create("random_numbers.txt")
	if err != nil {
		log.Error(err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err.Error())
		}
	}()
	nums := make(chan int, 2)
	startTime := time.Now()

	go addNumbers(nums)
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func() {
			writeToFile(f, nums)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Now().Sub(startTime))
}

func addNumbers(c chan<- int) {
	for i := 0; i < 100000; i++ {

		c <- rand.Int()
	}
	close(c)
}

func writeToFile(fpath *os.File, ch <-chan int) {
	for i := range ch {
		fmt.Fprintln(fpath, i)
		time.Sleep(30 * time.Microsecond)
	}
}
