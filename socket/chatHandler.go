package socket

import (
	"encoding/json"
	"log"

	"github.com/lwinmgmg/chat/models"
	"golang.org/x/net/websocket"
)

func (socketHandler *SocketHandler) HandleChat(uid string, mesg models.ChatData, ws *websocket.Conn) {
	mesg.Message.UserID = uid
	var err error
	switch mesg.ChatType {
	case models.CT_NEW:
		err = socketHandler.ChatTypeNew(uid, &mesg, ws)
	case models.CT_SEND:
		err = socketHandler.ChatTypeSend(uid, &mesg, ws)
	case models.CT_TYPE:
	}
	if err != nil {
		return
	}
	data, err := json.Marshal(mesg)
	if err != nil {
		log.Println("Error on json dump mesg", err, mesg)
		return
	}
	if mesg.ConversationID != 0 {
		convUserList, err := models.GetUidsByConversationID(mesg.ConversationID, PgDb)
		if err != nil {
			log.Println("Error", err)
			return
		}
		for i := 0; i < len(convUserList); i++ {
			if userInfo, ok := socketHandler.UserMap[convUserList[i].UserID]; ok {
				for _, wsConn := range userInfo.Conn {
					if _, err := wsConn.Write(data); err != nil {
						log.Println("Error on sending clients", err)
					}
				}
			}
		}
	}
}
