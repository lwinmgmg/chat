package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatType string

const (
	CT_SEND     ChatType = "send"
	CT_NEW      ChatType = "new"
	CT_TYPE     ChatType = "type"
	CT_UN_TYPE  ChatType = "untype"
	CT_REACT    ChatType = "react"
	CT_DEL      ChatType = "delete"
	CT_EDIT     ChatType = "edit"
	CT_C_CREATE ChatType = "ccreate"
)

type ChatMesg struct {
	ID            primitive.ObjectID `json:"id,omitempty"`
	ParentId      primitive.ObjectID `json:"parent_id,omitempty"`
	Message       string             `json:"mesg,omitempty"`
	AttachmentURL string             `json:"att_url,omitempty"`
	IsEdited      bool               `json:"is_edited"`
	Status        MessageStatus      `json:"status,omitempty"`
	UserID        string             `json:"user_id,omitempty"`
	UserList      []string           `json:"user_list,omitempty"`
	ReactionID    uint               `json:"react_id,omitempty"`
	UpdatedTime   uint               `json:"updated_time"`
	CreatedTime   uint               `json:"created_time"`
	UpdatedFields []string           `json:"fields,omitempty"`
}

type ChatData struct {
	Name             string           `json:"name,omitempty"`
	ConversationID   uint             `json:"conversation_id,omitempty"`
	ConversationType ConversationType `json:"conv_type,omitempty"`
	LastMesgID       string           `json:"last_mesg_id,omitempty"`
	ChatType         ChatType         `json:"chat_type,omitempty"`
	Message          ChatMesg         `json:"mesg,omitempty"`
	ImageURL         string           `json:"img_url,omitempty"`
}
