package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	username string
	ctime    string
}

func main() {
	token1 := jwt.New(jwt.SigningMethodHS256)
	claims := token1.Claims.(jwt.MapClaims)
	claims["username"] = "rohit"
	claims["date"] = time.Now().Format("2006-01-02 15:04:05")
	tokenString1, _ := token1.SignedString([]byte("helo"))

	claim := Claim{"Rohit", time.Now().Format("2006-01-02 15:04:05")}
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{"username": claim.username, "time": claim.ctime})
	tokenString2, _ := token2.SignedString([]byte("hello"))

	verifyToken(tokenString1)
	verifyToken(tokenString2)
}

func verifyToken(t string) {
	token, err := jwt.Parse(t, func(tok *jwt.Token) (interface{}, error) {
		fmt.Println(tok)
		_, ok := tok.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Invalid")
		}
		return []byte("hello"), nil
	})
	if err != nil {
		fmt.Println("Invalid")
	}
	if token.Valid {
		fmt.Println("")
	}
}
