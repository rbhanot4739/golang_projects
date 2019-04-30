package main

import (
	"fmt"
	"strings"
)

var days = []string{"Monday", "Tuesday", "Wednesday"}

func Joiner() {

	fmt.Println(strings.Join(days, "-"))
}
