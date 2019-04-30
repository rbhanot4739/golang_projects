package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	usr, _ := user.Current() // gives the current user running the program
	homeDir := usr.HomeDir   // gives the HomeDirectory of currently logged user
	file, err := os.Open(homeDir + "/.zshrc")
	if err != nil {
		panic("Error opening the file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	// readByChunks(file, 10)
	readByLines(file)
}

func readByChunks(f *os.File, n int) {
	reader := bufio.NewReader(f)
	chunk := make([]byte, n)
	for {
		_, err := reader.Read(chunk)
		if err != nil {
			fmt.Println("Error reading the file")
			break
		}
		fmt.Println(string(chunk))
	}
}

func readByLines(f *os.File) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
