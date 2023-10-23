package models

import (
	"context"
	"fmt"

	"github.com/lwinmgmg/chat/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

const (
	MongoDbName = "chat"
)

var (
	PgDb        *gorm.DB      = services.PgDb
	MongoClient *mongo.Client = services.GetMongoClient()
	MongoDb     *mongo.Database
)

func init() {
	conversation := &Conversation{}
	conversationUser := &ConversationUser{}
	if err := PgDb.AutoMigrate(conversation, conversationUser); err != nil {
		panic(err)
	}
	mesg := &Message{}
	MongoDb = MongoClient.Database(MongoDbName)
	mesgCol := MongoDb.Collection(mesg.GetCollection())
	if data, err := mesgCol.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "user_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "conversation_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "created_time", Value: -1}},
		},
		{
			Keys: bson.D{
				{Key: "created_time", Value: -1},
				{Key: "conversation_id", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "_id", Value: -1},
				{Key: "conversation_id", Value: 1},
			},
		},
	}); err != nil {
		panic(err)
	} else {
		fmt.Println(data)
	}
	fmt.Println("Done index")
}

func GetCollection(colName string, mongoDb *mongo.Database) *mongo.Collection {
	return mongoDb.Collection(colName)
}
