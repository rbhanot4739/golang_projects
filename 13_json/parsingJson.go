package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error(err.Error())
	}
	coll := client.Database("testdb").Collection("data")
	cursor, _ := coll.Find(ctx, bson.M{})
	for cursor.Next(ctx) {
		var record interface{}
		cursor.Decode(&record)
		fmt.Printf(record)
	}
}
