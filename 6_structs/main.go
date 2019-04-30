package main

import (
	"fmt"
)

type Books struct {
	title  string
	author string
	id     int
}

func main() {
	var book1 Books
	book1.author = "xyz"
	book1.id = 123
	book1.title = "Testing a book"

	book2 := Books{title: "This is the 2nd book", id: 788, author: "John Donovan"}

	fmt.Println(book1.title)
	fmt.Println(book2.title)
	passingStructs(book1)
	fmt.Println(book1.title)

	var ptr1 *Books
	ptr1 = &book1
	fmt.Println(ptr1.title)
	passingStructPtr(ptr1)
	fmt.Println(ptr1.title)

}

func passingStructs(x Books) {
	x.title = "New title"
	x.author = "New Author"
}

func passingStructPtr(x *Books) {
	x.title = "Changed title"
}
