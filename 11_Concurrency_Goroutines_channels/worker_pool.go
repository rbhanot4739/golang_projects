package main

import (
	"fmt"
	"sync"
	"time"
)

var jobs = make(chan int, 1)
var res = make(chan int, 1)
var wg sync.WaitGroup

func main() {
	st := time.Now()

	go producer(55)

	for j := 1; j < 5; j++ {
		wg.Add(1)
		go func() {
			workerFunc(jobs, res)
			wg.Done()
		}()
	}

	go consumer()
	wg.Wait()
	close(res)

	ed := time.Now()
	fmt.Println(ed.Sub(st))
}

func producer(n int) {
	for i := 0; i < n; i++ {
		jobs <- i
	}
	close(jobs)
}

func workerFunc(i <-chan int, o chan<- int) {
	for v := range i {
		o <- doSomething(v)
	}
}

func consumer() {
	for v := range res {
		fmt.Println(v)
	}
}

func doSomething(v int) int {
	first, second := 0, 1
	var third int
	for i := 1; i <= v; i++ {
		third = first + second
		first, second = second, third
	}
	time.Sleep(10 * time.Millisecond)
	return third

}
