package socket

import (
	"golang.org/x/net/websocket"
)

func (socket *SocketHandler) AuthFunc(ws *websocket.Conn) (string, error) {
	uidB, err := socket.ReadMesg(ws)
	if err != nil {
		return "", err
	}
	return string(uidB), nil
}
