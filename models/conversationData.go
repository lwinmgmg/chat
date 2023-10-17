package models

type ConversationData struct {
	ConvType ConversationType `json:"conv_type"`
	UserList []string         `json:"user_list"`
}
