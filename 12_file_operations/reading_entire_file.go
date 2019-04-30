package main

import (
	"fmt"
	"io/ioutil"
	"flag"
)

func main() {
	fptr := flag.String("fpath", "/apps/nttech/rbhanot/.zshrc", "provide the file path")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("Error reading the file")
	} else {
		fmt.Println("Byte slice of the file", data)
		fmt.Println("String contents of the file", string(data))
	}
}
