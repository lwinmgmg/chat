package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageStatus uint8

const (
	SENT     MessageStatus = 1
	RECEIVED MessageStatus = 2
	SEEN     MessageStatus = 3
)

type Message struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ParentID       primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id"`
	UserId         string             `bson:"user_id,omitempty" json:"user_id"`
	ExcludedUsers  []string           `bson:"excluded_users,omitempty" json:"excluded_users"`
	ConversationID uint               `bson:"conversation_id,omitempty" json:"conversation_id"`
	Message        string             `bson:"message" json:"mesg"`
	AttachmentURL  string             `bson:"attachment_url,omitempty" json:"att_url"`
	OldMessages    []Message          `bson:"old_messages" json:"old_messages"`
	Status         MessageStatus      `bson:"status" json:"status"`
	IsEdited       bool               `bson:"is_edited" json:"is_edited"`
	UpdatedTime    uint               `bson:"updated_time" json:"updated_time"`
	CreatedTime    uint               `bson:"created_time" json:"created_time"`
}

func (conv *Message) GetCollection() string {
	return "message"
}

func GetMessages(convId uint, lastMesgId string, limit int64, dest any, db *mongo.Database) error {
	mesg := Message{}
	col := GetCollection(mesg.GetCollection(), db)
	objId, err := primitive.ObjectIDFromHex(lastMesgId)
	if err != nil {
		return err
	}
	cur, err := col.Find(context.TODO(), bson.D{{Key: "$and",
		Value: bson.A{
			bson.D{{Key: "conversation_id", Value: convId}},
			bson.D{{Key: "_id", Value: bson.D{{Key: "$lte", Value: objId}}}},
		},
	}}, &options.FindOptions{
		Limit: &limit,
		Sort:  bson.D{{Key: "_id", Value: -1}},
	})
	if err != nil {
		return err
	}
	return cur.All(context.TODO(), dest)
}

func (mesg *Message) Create(db *mongo.Database) error {
	col := GetCollection(mesg.GetCollection(), db)
	inserted, err := col.InsertOne(context.Background(), mesg)
	if err != nil {
		return err
	}
	switch insertedId := inserted.InsertedID.(type) {
	case primitive.ObjectID:
		mesg.ID = insertedId
		return nil
	}
	return fmt.Errorf("err inserted id - %v", inserted.InsertedID)
}

func (mesg *Message) Delete(db *mongo.Database) error {
	col := GetCollection(mesg.GetCollection(), db)
	_, err := col.DeleteOne(context.TODO(), bson.D{{
		Key:   "_id",
		Value: mesg.ID,
	}})
	return err
}

func NewMessage(uid string, conId uint, mesg string, obj ...struct{ Attach string }) *Message {
	attachUrl := ""
	if len(obj) > 0 {
		attachUrl = obj[0].Attach
	}
	return &Message{
		UserId:         uid,
		ConversationID: conId,
		Message:        mesg,
		AttachmentURL:  attachUrl,
		UpdatedTime:    uint(time.Now().Unix()),
		CreatedTime:    uint(time.Now().Unix()),
	}
}
