package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client = nil
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	MongoClient, err = ConnectMongo(ctx)
	if err != nil {
		panic(fmt.Sprintf("Error on connecting mongo client : %v", err))
	}
}

func GetMongoClient() *mongo.Client {
	var err error
	if MongoClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		MongoClient, err = ConnectMongo(ctx)
		if err != nil {
			panic("Error on connection Mongo")
		}
	}
	return MongoClient
}

func ConnectMongo(ctx context.Context) (*mongo.Client, error) {
	var uri string = fmt.Sprintf(
		"mongodb://%v:%v@%v:%v",
		Env.Settings.Mongo.Login,
		Env.Settings.Mongo.Password,
		Env.Settings.Mongo.Host,
		Env.Settings.Mongo.Port,
	)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	return mongo.Connect(ctx, opts)
}
