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
	Excludes  map[string]bool   `bson:"excludes" json:"excludes"`
	Data      map[string]string `bson:"data" json:"data"`
}

func main() {
	fmt.Println("Main Starting")
	Connect()
}

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error(err.Error())
	}
	coll := client.Database("cfengine").Collection("configs")
	var hosts []HostConfig
	cur, err := coll.Find(ctx, bson.M{"type": "host"})
	for cur.Next(ctx) {
		var host HostConfig
		cur.Decode(&host)
		host, _ = ReadOne(coll, host.Hostname)
		hosts = append(hosts, host)
	}

	fmt.Printf("%v \n", hosts)
}

func ReadOne(collection *mongo.Collection, hostname string) (HostConfig, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var record HostConfig
	err := collection.FindOne(ctx, bson.M{"_id": hostname}).Decode(&record)
	//err := m.collection.FindOne(ctx, HostConfig{Hostname: host}).Decode(&record)
	if err != nil {
		log.Error(err.Error())
		return HostConfig{}, err
	}

	for _, val := range record.Groups {
		var grp GroupConfig
		collection.FindOne(ctx, bson.M{"type": "group", "_id": val}).Decode(&grp)
		for k, v := range grp.Data {
			if _, ok := record.Excludes[k]; !ok {

				if val, ok := record.Overrides[k]; ok {
					record.Data[k] = val
				} else {
					record.Data[k] = v
				}
			}
		}
	}

	return record, nil
}
