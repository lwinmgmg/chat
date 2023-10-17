package models

import "gorm.io/gorm"

type ConversationUser struct {
	gorm.Model
	Name           string
	UserID         string `gorm:"index;"`
	ConversationID uint   `gorm:"index;"`
	Conversation   Conversation
}

func (convUser *ConversationUser) Create(tx *gorm.DB) error {
	return tx.Create(convUser).Error
}

func GetUidsByConversationID(convId uint, tx *gorm.DB) ([]ConversationUser, error) {
	var convUserList []ConversationUser
	if err := tx.Model(&ConversationUser{}).Where(&ConversationUser{
		ConversationID: convId,
	}).Find(&convUserList).Error; err != nil {
		return nil, err
	}
	return convUserList, nil
}
