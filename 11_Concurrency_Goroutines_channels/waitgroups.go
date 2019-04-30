package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(j int, grp *sync.WaitGroup) {
	fmt.Println("Worker ", j)
	time.Sleep(47 * time.Millisecond)
	grp.Done()
}

func main() {
	fmt.Println("Started Main")
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
