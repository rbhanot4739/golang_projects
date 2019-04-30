package main

import (
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.google.com")
	bytes, _ := ioutil.ReadAll(resp.Body)
	text_data := string(bytes)
	fmt.Println(text_data)
	resp.Body.Close()
}
