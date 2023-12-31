package models

type ConversationData struct {
	Name     string           `json:"name"`
	Active   bool             `json:"active"`
	ConType  ConversationType `json:"conv_type"`
	UserID   string           `json:"user_id"`
	ImageURL string           `json:"img_url"`
}

type ConversationCreateData struct {
	ConvType ConversationType `json:"conv_type"`
	UserList []string         `json:"user_list"`
}

type ConversationInfo struct {
	ID               uint             `json:"id"`
	ConversationType ConversationType `json:"conv_type"`
	LastMesgID       string           `json:"last_mesg_id"`
}

type ConversationDetail struct {
	ID               uint             `json:"id"`
	ConversationType ConversationType `json:"conv_type"`
	Name             string           `json:"name"`
	Active           bool             `json:"active"`
	UserID           string           `json:"user_id"`
	ImageURL         string           `json:"img_url"`
	LastMesgId       string           `json:"last_mesg_id"`
	ConvUsers        []ConvUserDetail `json:"conv_users"`
}
