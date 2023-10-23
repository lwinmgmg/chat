package socket

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lwinmgmg/chat/models"
	"github.com/lwinmgmg/chat/services"
	"github.com/lwinmgmg/chat/utils"
	gmodel "github.com/lwinmgmg/gmodels/golang/models"
	"golang.org/x/net/websocket"
)

type HandleType uint

const (
	HandleAdd   HandleType = 1
	HandleClose HandleType = 2
)

var (
	PgDb    = services.PgDb
	MongoDb = models.MongoDb
)

type UserInfo struct {
	User *gmodel.User
	Conn map[string]*websocket.Conn
}

type HandleConn struct {
	Ws         *websocket.Conn
	UID        *string
	UuidCode   string
	User       *gmodel.User
	HandleType HandleType
}

type SocketHandler struct {
	Auth       bool
	ConnAge    time.Duration
	UserMap    map[string]*UserInfo
	HandleConn chan HandleConn
	CloseChan  chan struct{}
}

func (socketHandler *SocketHandler) Init() {
	socketHandler.HandleConn = make(chan HandleConn, 100)
	socketHandler.CloseChan = make(chan struct{})
	go func(ch <-chan HandleConn) {
		stop := false
		for !stop {
			select {
			case connData := <-socketHandler.HandleConn:
				switch connData.HandleType {
				case HandleClose:
					connData.Ws.Close()
					if *connData.UID != "" {
						if len(socketHandler.UserMap[*connData.UID].Conn) > 1 {
							delete(socketHandler.UserMap[*connData.UID].Conn, connData.UuidCode)
						} else {
							delete(socketHandler.UserMap, *connData.UID)
						}
					}
				case HandleAdd:
					userInfo, ok := socketHandler.UserMap[*connData.UID]
					if ok {
						userInfo.Conn[connData.UuidCode] = connData.Ws
					} else {
						var newUserMap map[string]*websocket.Conn = make(map[string]*websocket.Conn, 2)
						newUserMap[connData.UuidCode] = connData.Ws
						socketHandler.UserMap[*connData.UID] = &UserInfo{
							User: connData.User,
							Conn: newUserMap,
						}
					}
				}
			case <-socketHandler.CloseChan:
				stop = true
			}
		}
	}(socketHandler.HandleConn)
}

func (socketHandler *SocketHandler) Close() {
	socketHandler.CloseChan <- struct{}{}
}

func (socketHandler *SocketHandler) HandleSocket(ws *websocket.Conn) {
	uid := ""
	uuidCode := uuid.New().String()
	var err error
	// defer section
	defer func(c *websocket.Conn) {
		log.Println("Close the connection", ws.RemoteAddr())
		socketHandler.HandleConn <- HandleConn{
			Ws:         ws,
			UID:        &uid,
			UuidCode:   uuidCode,
			HandleType: HandleClose,
			User:       nil,
		}
	}(ws)
	log.Println("Got one connection", ws.RemoteAddr(), ws.LocalAddr(), ws.Request().Header["Sec-Websocket-Key"])

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
	log.Println("Authenticated as", user.Code, uuidCode)
	// ws.Write([]byte("Hello" + user.Partner.FirstName + " " + user.Partner.LastName))
	uid = user.Code

	// Assigning User Map
	socketHandler.HandleConn <- HandleConn{
		Ws:         ws,
		UID:        &uid,
		UuidCode:   uuidCode,
		HandleType: HandleAdd,
		User:       user,
	}

	for {
		mesgB, err := socketHandler.ReadMesg(ws)
		if err != nil {

			if errors.Is(err, io.EOF) {
				break
			}
			log.Println("Error on reading message", err)
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
		default:
			fmt.Println("Default", string(mesgB))
		}
	}
}
