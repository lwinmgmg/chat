package socket

import "golang.org/x/net/websocket"

func (socket *SocketHandler) CallBack(uid *string, ws *websocket.Conn) {
	ws.Close()
	if *uid != "" {
		delete(socket.UserMap, *uid)
	}
}
