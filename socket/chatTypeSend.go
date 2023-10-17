package socket

import (
	"time"

	"github.com/lwinmgmg/chat/models"
	"golang.org/x/net/websocket"
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
	switch data.ConversationType {
	case models.NormalCon:
		if err := validateNormalConSend(data); err != nil {
			return err
		}
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
		data.Message.Status = mesg.Status
		data.Message.ID = mesg.ID
	case models.GroupCon:
		if err := validateGroupConSend(data); err != nil {
			return err
		}
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
		data.Message.Status = mesg.Status
		data.Message.ID = mesg.ID
	}
	return nil
}
