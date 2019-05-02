package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Main GORoutine")
	f, err := os.OpenFile("random_numbers.txt", os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		log.Error(err.Error())
	}
	

}
