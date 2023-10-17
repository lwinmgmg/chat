package socket

import (
	"github.com/lwinmgmg/chat/grpc/client"
	gmodels "github.com/lwinmgmg/gmodels/golang/models"
	"golang.org/x/net/websocket"
)

func (socketHandler *SocketHandler) AuthFunc(ws *websocket.Conn) (*gmodels.User, error) {
	tokenB, err := socketHandler.ReadMesg(ws)
	if err != nil {
		return nil, err
	}
	user, err := client.GetUserByToken(string(tokenB))
	return user, err
}
