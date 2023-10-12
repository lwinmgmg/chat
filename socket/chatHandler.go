package socket

import (
	"encoding/json"
	"fmt"

	"github.com/lwinmgmg/chat/models"
	"golang.org/x/net/websocket"
)

func (socket *SocketHandler) HandleChat(uid string, mesg models.ChatData, ws *websocket.Conn) {
	switch mesg.ChatType {
	case models.CT_NEW:
		fmt.Println("NEW", mesg.Message)
		data, _ := json.Marshal(mesg)
		ws.Write(data)
	case models.CT_TYPE:
		fmt.Println("NEW", mesg.Message)
		data, _ := json.Marshal(mesg)
		ws.Write(data)
	case models.CT_SEND:
		fmt.Println("SEND", mesg.Message)
	}
}
