package main

import "fmt"

func main() {
	valid := []interface{}{"a", "b", "c"}
	invalid := []interface{}{}
	hostgroups := []interface{}{"b", "d", "f"}
	// var final []interface{}

	for i, val := range valid {
		if !SliceContains(hostgroups, val) {
			invalid = append(invalid, val)
			valid = append(valid[:i], valid[i+1:]...)
		}
	}
	// valid = final
	// final = nil
	// delete final
	fmt.Printf("%v -- %p\n", valid, &valid)

	fmt.Println(invalid)
	fmt.Println(hostgroups)
}

func SliceContains(a []interface{}, v interface{}) bool {
	for _, val := range a {
		if v == val {
			return true
		}
	}
	return false
}
