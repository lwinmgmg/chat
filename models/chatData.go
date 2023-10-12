package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatType string

const (
	CT_SEND     ChatType = "sent"
	CT_NEW      ChatType = "new"
	CT_TYPE     ChatType = "type"
	CT_UN_TYPE  ChatType = "untype"
	CT_REACT    ChatType = "react"
	CT_DEL      ChatType = "delete"
	CT_EDIT     ChatType = "edit"
	CT_C_CREATE ChatType = "ccreate"
)

type ChatMesg struct {
	ID                primitive.ObjectID `json:"id"`
	Message           string             `json:"mesg"`
	AttachmentURL     string             `json:"att_url"`
	ConversationID    uint               `json:"conv_id"`
	ConversationType  ConversationType   `json:"conv_type"`
	ConversationName  string             `json:"conv_name"`
	ConversationImage string             `json:"conv_img"`
	UserID            string             `json:"uid"`
	UserList          []string           `json:"user_list"`
	ReactionID        uint               `json:"react_id"`
	UpdatedFields     []string           `json:"fields"`
}

type ChatData struct {
	ConversationID uint     `json:"cid"`
	ChatType       ChatType `json:"chat_type"`
	Message        ChatMesg `json:"mesg"`
}
