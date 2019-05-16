package main

import (
	"fmt"
	whatever "rbhanot/Pack1"
	helloworld "rbhanot/Pack2"
)

func main() {
	fmt.Println(whatever.Greet("Rohit"))
	whatever.Bye("Rohit")
	helloworld.HelloWorld()
}
