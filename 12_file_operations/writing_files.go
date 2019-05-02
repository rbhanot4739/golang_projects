package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Writing Files")
	f, err := os.Create("data.txt")

	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err.Error())
		}
	}()

	if err != nil {
		log.Error(err.Error())
		return
	}
	// Writing string to file
	l, _ := f.WriteString("Writing the string 'Hello' to the file\n")
	fmt.Println(l, "Bytes written to the file.")

	// Writing byte slice to file
	l, _ = f.Write([]byte("Hello\n"))
	fmt.Println(l, "Bytes written to the file.")

	// Writing multiple lines to file
	lines := []string{"This is line 1", "Line 2", "Line3"}
	for _, line := range lines {
		fmt.Fprintln(f, line)
	}
}
