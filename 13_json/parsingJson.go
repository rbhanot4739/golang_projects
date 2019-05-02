package main

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// p1 := Person{"Mike", "Smith", 32}
	// a1 := Addr{"Ggn", "Haryana"}
	// arr := []interface{}{p1, a1}

	// d, _ := json.Marshal(arr)

	// fmt.Println(string(d))
	Connect()
}

type GroupConfig struct {
	Groupname string   `bson:"_id" json:"groupname"`
	Type      string   `bson:"type" json:"type"`
	Tags      []string `bson:"tags" json:"tags"`
}

type HostConfig struct {
	Hostname  string            `bson:"_id" json:"hostname"`
	Type      string            `bson:"type" json:"type"`
	Groups    []string          `bson:"groups" json:"groups"`
	Overrides map[string]string `bson:"overrides" json:"overrides"`
	Excludes  []string          `bson:"excludes" json:"excludes"`
	tags      interface{}       `bson:tags json:tags`
}

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error(err.Error())
	}
	coll := client.Database("cfengine").Collection("configs")
	opts := options.FindOne()
	projection := bson.M{"tags": 1, "_id": 0}
	opts.SetProjection(projection)
	var record HostConfig
	err = coll.FindOne(ctx, bson.M{"_id": "install237"}).Decode(&record)

	// cursor, _ := coll.Find(ctx, bson.M{"type": "host"})
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	var record HostConfig
	// 	cursor.Decode(&record)
	// 	// for every group host is part of, we need to get tag information from that group
	for _, val := range record.Groups {
		var grps interface{}
		coll.FindOne(ctx, bson.M{"type": "group", "_id": val}, opts).Decode(&grps)
		record.tags = grps
		// data, _ := json.Marshal(record)
		fmt.Println(record)
	}
	// }
}
