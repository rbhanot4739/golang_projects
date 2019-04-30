package main

import (
	"fmt"
)

func main() {
	i := 10
	var j *int // declares j as a integer pointer
	j = &i
	fmt.Println(j, i)
	k := 10.9
	l := &k
	fmt.Println(l, k)
	var p **int
	var q **float64
	p = &j
	q = &l
	fmt.Println(p, q, **p, **q, *p, *q)
}
