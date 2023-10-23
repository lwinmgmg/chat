package socket

import (
	"time"

	"github.com/lwinmgmg/chat/models"
	"golang.org/x/net/websocket"
	"gorm.io/gorm"
)

func commonValidateSend(data *models.ChatData) error {
	return nil
}

func validateNormalConSend(data *models.ChatData) error {
	if err := commonValidateSend(data); err != nil {
		return err
	}
	return nil
}

func validateGroupConSend(data *models.ChatData) error {
	if err := commonValidateSend(data); err != nil {
		return err
	}
	return nil
}

func (socketHandler *SocketHandler) ChatTypeSend(uid string, data *models.ChatData, ws *websocket.Conn) error {
	data.Message.UserID = uid
	return PgDb.Transaction(
		func(tx *gorm.DB) error {
			switch data.ConversationType {
			case models.NormalCon:
				if err := validateNormalConSend(data); err != nil {
					return err
				}
			case models.GroupCon:
				if err := validateGroupConSend(data); err != nil {
					return err
				}
			}
			mesg := models.Message{
				ParentID:       data.Message.ParentId,
				UserId:         data.Message.UserID,
				ConversationID: data.ConversationID,
				Message:        data.Message.Message,
				AttachmentURL:  data.Message.AttachmentURL,
				Status:         models.SENT,
				IsEdited:       false,
				UpdatedTime:    uint(time.Now().UTC().Unix()),
				CreatedTime:    uint(time.Now().UTC().Unix()),
			}
			if err := mesg.Create(MongoDb); err != nil {
				return err
			}
			data.Message.Status = mesg.Status
			data.Message.ID = mesg.ID
			data.Message.UpdatedTime = mesg.UpdatedTime
			data.Message.CreatedTime = mesg.CreatedTime
			data.LastMesgID = mesg.ID.Hex()
			return models.UpdateLastMesgId(data.ConversationID, data.LastMesgID, tx)
		},
	)
}
