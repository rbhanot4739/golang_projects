package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Fname string
	Lname string
	Age   int
}

type Addr struct {
	City  string
	State string
}

func main() {
	p1 := Person{"Mike", "Smith", 32}
	a1 := Addr{"Ggn", "Haryana"}
	arr := []interface{}{p1, a1}

	d, _ := json.Marshal(arr)

	fmt.Println(string(d))
}
