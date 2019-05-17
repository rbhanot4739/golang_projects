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
	verifyToken("jdsjdsjsdhsdjhsdjhsjsdhsdjfh")
}

func verifyToken(t string) {
	token, err := jwt.Parse(t, func(tok *jwt.Token) (interface{}, error) {
		if tok.Header["method"] == "RSA" {
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
			if _, ok := tok.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unknown Signing method")
			}
			return validationKey, nil
		}
		_, ok := tok.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unknown Signing method")
		}
		return []byte("hello"), nil

	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Fatal("That's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				log.Fatal("Expired token")
			} else {
				log.Fatal("Couldn't handle this token:", err)
			}
		} else {
			log.Fatal("Couldn't handle this token:", err)
		}
	}
	if token.Valid {
		fmt.Println("Token Validated")
		fmt.Println("Claims are")
		fmt.Println(token.Claims.(jwt.MapClaims))
	}
}




/ GenerateJWToken = Generates the JWT
func GenerateJWToken(uid string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["time"] = time.Now().Unix()

	signingKey, err := createSigningKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	return tokenString, nil
}

// VerifyToken - VErfies JWT
func VerifyToken(t string) (bool, error) {
	token, err := jwt.Parse(t, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Invalid Signing method")
		}
		createVerifyKey, err := createVerifyKey()
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		return createVerifyKey, nil
	})
	if err != nil {
		var msg string
		// type assertion
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				msg = "Malformed Token: " + err.Error()
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				msg = "Token expired or Inactive: " + err.Error()
			} else {
				msg = "Could not handle token: " + err.Error()

			}
			return false, fmt.Errorf(msg)
		}
	}
	if token.Valid {
		return true, nil
	}
	return false, fmt.Errorf("Invalid token")
}

func createSigningKey() (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(os.Getenv("APP_PRIV_KEY"))
	if err != nil {
		return nil, fmt.Errorf("Unable to open the private key file %v", err.Error())
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key %v", err.Error())
	}

	return key, nil

}

func createVerifyKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(os.Getenv("APP_PUB_KEY"))
	if err != nil {
		return nil, fmt.Errorf("Unable to open the public key file %v", err.Error())
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse public key %v", err.Error())
	}
	return key, nil

}

func main() {
	ts, _ := GenerateJWToken("Roht")
	if ok, _ := VerifyToken(ts); !ok {
		fmt.Println("Token not verified")
	} else {
		fmt.Println("Verified")
	}

}

