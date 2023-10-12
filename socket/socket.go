package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lwinmgmg/chat/models"
	"github.com/lwinmgmg/chat/utils"
	"golang.org/x/net/websocket"
)

type SocketHandler struct {
	Auth    bool
	ConnAge time.Duration
	UserMap map[string]*websocket.Conn
}

func (socket *SocketHandler) HandleSocket(ws *websocket.Conn) {
	uid := ""
	var err error
	// defer section
	defer func(c *websocket.Conn) {
		log.Println("Close the connection", ws.RemoteAddr())
		socket.CallBack(&uid, c)
	}(ws)
	log.Println("Got one connection", ws.RemoteAddr())

	// Setting connection life
	// if err := ws.SetDeadline(time.Now().UTC().Add(socket.ConnAge)); err != nil {
	// 	log.Println("Error on setting socket deadline")
	// 	return
	// }

	// Authentication
	uid, err = socket.AuthFunc(ws)
	if err != nil {
		log.Println("Error on auth", err)
	}
	ws.Write([]byte("Hello"))
	// Assigning User Map
	socket.UserMap[uid] = ws

	for {
		mesgB, err := socket.ReadMesg(ws)
		if err != nil {
			log.Println(err)
			break
		}
		var mesg models.SocketData
		if err := json.Unmarshal(mesgB, &mesg); err != nil {
			log.Println("Error mesg format :", err)
			continue
		}
		switch mesg.SocketType {
		case models.CHAT:
			var data models.ChatData
			utils.MapToStruct[any](mesg.Data, &data)
			socket.HandleChat(uid, data, ws)
			break
		default:
			fmt.Println("Default", string(mesgB))
			break
		}
	}
}
