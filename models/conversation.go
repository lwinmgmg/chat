package models

import "gorm.io/gorm"

type ConversationType uint

const (
	NormalCon ConversationType = 1
	GroupCon  ConversationType = 2
)

type Conversation struct {
	gorm.Model
	Name     string `gorm:"size:40"`
	Active   bool
	ConType  ConversationType `gorm:"index"`
	UserID   string           `gorm:"index"`
	ImageURL string           `gorm:"size: 256"`
}

func (conv *Conversation) GetCollection() string {
	return "conversation"
}
