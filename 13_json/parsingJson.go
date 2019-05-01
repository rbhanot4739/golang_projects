package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Person struct {
	Fname string
	Lname string
	Age   int
}

type Addr struct {
	City  string
	State string
}

func main() {
	p1 := Person{"Mike", "Smith", 32}
	a1 := Addr{"Ggn", "Haryana"}
	arr := []interface{}{p1, a1}

	d, _ := json.Marshal(arr)

	fmt.Println(string(d))
	Connect()
}

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Error(err.Error())
	}
	coll := client.Database("testdb").Collection("data")
	cursor, _ := coll.Find(ctx, bson{})
}
