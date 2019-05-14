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

type GroupConfig struct {
	Groupname string            `bson:"_id" json:"groupname"`
	Type      string            `bson:"type" json:"type"`
	Data      map[string]string `bson:"data" json:"data"`
}

type HostConfig struct {
	Hostname  string            `bson:"_id" json:"hostname"`
	Type      string            `bson:"type" json:"type"`
	Groups    []string          `bson:"groups" json:"groups"`
	Overrides map[string]string `bson:"overrides" json:"overrides"`
	Excludes  []interface{}     `bson:"excludes" json:"excludes"`
	Data      map[string]string `bson:"data" json:"data"`
}

func main() {
	fmt.Println("Main Starting")
	Connect()
}

// Connect -
func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error(err.Error())
	}
	coll := client.Database("cfengine").Collection("configs")
	var host HostConfig
	projection := options.FindOne()
	projection.SetProjection(bson.M{"groups": 1})
	err = coll.FindOne(ctx, bson.M{"_id": "test800"}, projection).Decode(&host)

	fmt.Printf("%v \n", host.Groups)
}
