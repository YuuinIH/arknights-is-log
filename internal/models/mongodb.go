package models

import (
	"context"
	"fmt"

	"github.com/YuuinIH/arknights-is-log/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db *mongo.Client
)

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DATABASE.URI))
	if err != nil {
		panic(err)
	}
	/*defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()*/
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
	db = client
}
