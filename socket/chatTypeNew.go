package socket

import (
	"errors"
	"time"

	"github.com/lwinmgmg/chat/models"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
)

var (
	InvalidChatTypeSend = errors.New("Invalid data [ChatTypeSend]")
)

func validateNormalCon(data *models.ChatData) error {
	if len(data.Message.UserList) != 1 {
		return InvalidChatTypeSend
	}
	return nil
}

func validateGroupCon(data *models.ChatData) error {
	return nil
}

func (socketHandler *SocketHandler) ChatTypeNew(uid string, data *models.ChatData, ws *websocket.Conn) error {
	switch data.ConversationType {
	case models.NormalCon:
		if err := validateNormalCon(data); err != nil {
			return err
		}
		return PgDb.Transaction(func(tx *gorm.DB) error {
			conv, err := models.GetNormalConversation(uid, data.Message.UserList[0], tx)
			if err != nil {
				return err
			}
			data.ConversationID = conv.ID
			if data.Message.Message != "" || data.Message.AttachmentURL != "" {
				mesg := models.Message{
					ParentID:       data.Message.ParentId,
					UserId:         data.Message.UserID,
					ConversationID: data.ConversationID,
					Message:        data.Message.Message,
					AttachmentURL:  data.Message.AttachmentURL,
					Status:         models.SENT,
					UpdatedTime:    uint(time.Now().UTC().Unix()),
					CreatedTime:    uint(time.Now().UTC().Unix()),
				}
				if err := mesg.Create(MongoDb); err != nil {
					return err
				}
				data.Message.ID = mesg.ID
			}
			return nil
		})
	case models.GroupCon:
		if err := validateGroupCon(data); err != nil {
			return err
		}
		return PgDb.Transaction(func(tx *gorm.DB) error {
			conv := models.Conversation{
				Name:     data.Name,
				UserID:   uid,
				Active:   true,
				ConType:  models.GroupCon,
				ImageURL: data.ImageURL,
			}
			if err := conv.Create(tx); err != nil {
				return err
			}
			data.ConversationID = conv.ID
			if data.Message.Message != "" || data.Message.AttachmentURL != "" {
				mesg := models.Message{
					ParentID:       data.Message.ParentId,
					UserId:         data.Message.UserID,
					ConversationID: data.ConversationID,
					Message:        data.Message.Message,
					AttachmentURL:  data.Message.AttachmentURL,
					Status:         models.SENT,
					UpdatedTime:    uint(time.Now().UTC().Unix()),
					CreatedTime:    uint(time.Now().UTC().Unix()),
				}
				if err := mesg.Create(MongoDb); err != nil {
					return err
				}
				data.Message.ID = mesg.ID
			}
			return nil
		})
	}
	return nil
}
