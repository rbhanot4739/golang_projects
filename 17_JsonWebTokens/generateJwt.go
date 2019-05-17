package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	username string
	ctime    string
}

func main() {
	token1 := jwt.New(jwt.SigningMethodRS256)
	token1.Header["method"] = "RSA"
	claims := token1.Claims.(jwt.MapClaims)
	claims["username"] = "rohit"
	claims["date"] = time.Now().Unix()
	rsaKey, err := ioutil.ReadFile("rsa")
	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM(rsaKey)
	if err != nil {
		log.Fatal("Error parsing Private key")
	}

	tokenString1, err := token1.SignedString(signingKey)
	if err != nil {
		log.Fatal("Error generating token string")
	}

	claim := Claim{"Rohit", time.Now().Format("2006-01-02 15:04:05")}
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{"username": claim.username, "time": claim.ctime})
	tokenString2, _ := token2.SignedString([]byte("hello"))

	verifyToken(tokenString1)
	verifyToken(tokenString2)
}

func verifyToken(t string) {
	token, err := jwt.Parse(t, func(tok *jwt.Token) (interface{}, error) {
		signingMethod := tok.Header["method"]
		// the public key need to converted to PkCS8 format
		// ssh-keygen -f id_rsa.pub -e -m pkcs8 >> rsa_pub.pem
		rsaPubKey, err := ioutil.ReadFile("rsa_pub.pem")
		if err != nil {
			log.Fatal(err.Error())
		}

		validationKey, err := jwt.ParseRSAPublicKeyFromPEM(rsaPubKey)
		if err != nil {
			log.Fatal("Error parsing public key", err.Error())
		}
		if signingMethod == "RSA" {
			if _, ok := tok.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unknown Signing method")
			}
			return validationKey, nil
		} else {
			_, ok := tok.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unknown Signing method")
			}
			return []byte("hello"), nil
		}
	})
	if err != nil {
		log.Fatal("Token not validated")
	}
	if token.Valid {
		fmt.Println("Valid")
	}
}
