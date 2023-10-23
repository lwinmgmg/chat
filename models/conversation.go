package models

import (
	"errors"

	"gorm.io/gorm"
)

type ConversationType uint

const (
	NormalCon ConversationType = 1
	GroupCon  ConversationType = 2
)

var (
	ErrConvNotFound = errors.New("conversation not found")
)

type Conversation struct {
	gorm.Model
	Name       string           `gorm:"size:40" json:"name"`
	Active     bool             `json:"active"`
	ConType    ConversationType `gorm:"index" json:"conv_type"`
	UserID     string           `gorm:"index" json:"user_id"`
	ImageURL   string           `gorm:"size: 256" json:"img_url"`
	LastMesgID string           `gorm:"index; size: 32" json:"last_mesg_id"`
}

func (conv *Conversation) GetCollection() string {
	return "conversation"
}

func (conv *Conversation) Create(tx *gorm.DB) error {
	return tx.Create(conv).Error
}

func (conv *Conversation) Get(convId uint, tx *gorm.DB) error {
	return tx.Model(conv).First(conv, convId).Error
}

func (conv *Conversation) SetActive(tx *gorm.DB) error {
	return tx.Model(conv).Update("active", true).Error
}

func (conv *Conversation) GetConvByUserId(uid string, dest any, tx *gorm.DB) error {
	return tx.Model(conv).Distinct("conversations.*").Joins("INNER JOIN conversation_users ON conversation_users.conversation_id=conversations.id AND conversation_users.user_id = ?", uid).Order("conversations.last_mesg_id DESC, conversations.updated_at DESC").Find(dest).Error
}

func UpdateLastMesgId(id uint, mesgId string, tx *gorm.DB) error {
	conv := &Conversation{}
	if err := tx.Model(conv).First(conv, id).Error; err != nil {
		return err
	}
	conv.LastMesgID = mesgId
	return tx.Save(conv).Error
}

func CreateNewNormalConv(uid1, uid2 string, tx *gorm.DB) (*Conversation, *ConversationUser, *ConversationUser, error) {
	var conv = &Conversation{
		Active:  false,
		ConType: NormalCon,
		UserID:  uid1,
	}
	if err := conv.Create(tx); err != nil {
		return nil, nil, nil, err
	}
	var convUser1 = &ConversationUser{
		UserID:         uid1,
		ConversationID: conv.ID,
	}
	if err := convUser1.Create(tx); err != nil {
		return nil, nil, nil, err
	}
	var convUser2 = &ConversationUser{
		UserID:         uid2,
		ConversationID: conv.ID,
	}
	if err := convUser2.Create(tx); err != nil {
		return nil, nil, nil, err
	}
	return conv, convUser1, convUser2, nil
}

func FindNormalConversation(uid1, uid2 string, tx *gorm.DB) (*Conversation, error) {
	var conv Conversation
	if err := tx.Raw(`
	SELECT DISTINCT ON (conv.id) conv.* FROM conversations AS conv
	INNER JOIN conversation_users AS conv_user ON conv.id=conv_user.conversation_id
	INNER JOIN conversation_users AS conv_user1 ON conv.id=conv_user1.conversation_id AND conv_user1.id!=conv_user.id
	WHERE conv.con_type=$1
	AND conv_user.user_id=$2
	AND conv_user1.user_id=$3
	LIMIT 1
	`, NormalCon, uid1, uid2).Scan(&conv).Error; err != nil {
		return &conv, err
	}
	if conv.ID == 0 {
		return &conv, ErrConvNotFound
	}
	return &conv, nil
}

func GetNormalConversation(uid1, uid2 string, tx *gorm.DB) (*Conversation, error) {
	conv, err := FindNormalConversation(uid1, uid2, tx)
	if err == ErrConvNotFound {
		conv, _, _, err := CreateNewNormalConv(uid1, uid2, tx)
		if err != nil {
			return conv, err
		}
		return conv, nil
	}
	return conv, err
}

func CreateNewGroupConv(uid string, userList []string, tx *gorm.DB) (*Conversation, error) {
	var conv = &Conversation{
		Active:  true,
		ConType: GroupCon,
		UserID:  uid,
	}
	if err := conv.Create(tx); err != nil {
		return nil, err
	}
	var convUser = &ConversationUser{
		UserID:         uid,
		ConversationID: conv.ID,
	}
	if err := convUser.Create(tx); err != nil {
		return nil, err
	}
	for _, user := range userList {
		var convUser = &ConversationUser{
			UserID:         user,
			ConversationID: conv.ID,
		}
		if err := convUser.Create(tx); err != nil {
			return nil, err
		}
	}
	return conv, nil
}
