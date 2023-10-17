package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lwinmgmg/chat/models"
	"github.com/lwinmgmg/chat/services"
	"github.com/lwinmgmg/chat/utils"
	gmodel "github.com/lwinmgmg/gmodels/golang/models"
	"golang.org/x/net/websocket"
)

var (
	PgDb    = services.PgDb
	MongoDb = models.MongoDb
)

type UserInfo struct {
	User *gmodel.User
	Conn *websocket.Conn
}

type SocketHandler struct {
	Auth    bool
	ConnAge time.Duration
	UserMap map[string]UserInfo
}

func (socketHandler *SocketHandler) HandleSocket(ws *websocket.Conn) {
	uid := ""
	var err error
	// defer section
	defer func(c *websocket.Conn) {
		log.Println("Close the connection", ws.RemoteAddr())
		socketHandler.CallBack(&uid, c)
	}(ws)
	log.Println("Got one connection", ws.RemoteAddr())

	// Setting connection life
	if err := ws.SetDeadline(time.Now().UTC().Add(socketHandler.ConnAge)); err != nil {
		log.Println("Error on setting socket deadline")
		return
	}

	// Authentication
	user, err := socketHandler.AuthFunc(ws)
	if err != nil {
		log.Println("Error on auth", err)
		return
	}
	ws.Write([]byte("Hello" + user.Partner.FirstName + " " + user.Partner.LastName))
	uid = user.Code
	// Assigning User Map
	socketHandler.UserMap[uid] = UserInfo{
		User: user,
		Conn: ws,
	}

	for {
		mesgB, err := socketHandler.ReadMesg(ws)
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
			socketHandler.HandleChat(uid, data, ws)
			break
		default:
			fmt.Println("Default", string(mesgB))
			break
		}
	}
}
