package socket

import "golang.org/x/net/websocket"

func (socketHandler *SocketHandler) CallBack(uid *string, ws *websocket.Conn) {
	ws.Close()
	if *uid != "" {
		delete(socketHandler.UserMap, *uid)
	}
}
