package models

type ConvUserData struct {
	UserID         string `json:"user_id"`
	Name           string `json:"name,omitempty"`
	ConversationID string `json:"conv_id"`
	LastMesgId     string `json:"last_mesg_id"`
}

type ConvUserDetail struct {
	Name           string `json:"name"`
	UserID         string `json:"user_id"`
	ConversationID string `json:"conv_id"`
}
