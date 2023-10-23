package socket

import "golang.org/x/net/websocket"

func (socketHandler *SocketHandler) CallBack(uid *string, uuidCode *string, ws *websocket.Conn) {
	ws.Close()
	if *uid != "" {
		if len(socketHandler.UserMap[*uid].Conn) > 1 {
			delete(socketHandler.UserMap[*uid].Conn, *uuidCode)
		} else {
			delete(socketHandler.UserMap, *uid)
		}
	}
}
