package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageStatus uint8

const (
	SENT     MessageStatus = 1
	RECEIVED MessageStatus = 2
	SEEN     MessageStatus = 3
)

type Message struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ParentID       primitive.ObjectID `bson:"parent_id,omitempty"`
	UserId         string             `bson:"user_id,omitempty"`
	ConversationID uint               `bson:"conversation_id,omitempty"`
	Message        string             `bson:"message"`
	AttachmentURL  string             `bson:"attachment_url,omitempty"`
	OldMessages    []Message          `bson:"old_messages"`
	Status         MessageStatus      `bson:"status"`
	IsEdited       bool               `bson:"is_edited"`
	UpdatedTime    uint               `bson:"updated_time"`
	CreatedTime    uint               `bson:"created_time"`
}

func (conv *Message) GetCollection() string {
	return "message"
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
	return fmt.Errorf("Error inserted ID %v", inserted.InsertedID)
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
