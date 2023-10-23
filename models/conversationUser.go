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

func (convUser *ConversationUser) GetByConvId(convId uint, dest any, tx *gorm.DB) error {
	if err := tx.Model(convUser).Where(&ConversationUser{
		ConversationID: convId,
	}).Find(dest).Error; err != nil {
		return err
	}
	return nil
}

func (conv *ConversationUser) GetConversationsByUserId(uid string, dest any, tx *gorm.DB) error {
	return tx.Exec("SELECT cu.conversation_id FROM conversation_users WHERE user_id=$1", uid).Error
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
