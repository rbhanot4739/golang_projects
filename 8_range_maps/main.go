package main

import "fmt"

func main() {

	arr := []string{"One", "Two", "Three"}

	for idx, val := range arr {
		fmt.Println(idx, val)
	}

	for idx, val := range "Hello" {
		fmt.Printf("Index--> %d, Value--> %c, Ascii Value--> %d\n", idx, val, val)
	}

	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	fmt.Println(countryCapitalMap)
	countryCapitalMap["Italy"] = "Venice"
	delete(countryCapitalMap, "Japan")
	fmt.Println(countryCapitalMap)
	// for - range
	for key, val := range countryCapitalMap {
		fmt.Printf("%s ===> %s\n", key, val)
	}

}
