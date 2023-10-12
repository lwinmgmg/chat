package models

import "gorm.io/gorm"

type ConversationUser struct {
	gorm.Model
	Name           string
	UserID         string `gorm:"index;"`
	ConversationID int    `gorm:"index;"`
	Conversation   Conversation
}
