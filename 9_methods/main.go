package main

import (
	"fmt"
)

type User struct {
	fname, lname string
	age          int
}

func (u User) makeEmail() string {
	return fmt.Sprintf("Email is %s%s@test.com", u.fname, u.lname)
}

func (uptr *User) birthday() {
	uptr.age += 1
}

func main() {
	user1 := User{"Mike", "Harley", 30}
	fmt.Println(user1.makeEmail())
	fmt.Println("User1 before B'day", user1)
	user1.birthday()
	fmt.Println("User1 before B'day", user1)
}
