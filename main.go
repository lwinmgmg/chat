package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lwinmgmg/chat/socket"
	"golang.org/x/net/websocket"
)

// This example demonstrates a trivial echo server.
func main() {
	pubSocket := socket.SocketHandler{
		Auth:    false,
		UserMap: make(map[string]*websocket.Conn, 100),
		ConnAge: time.Hour,
	}
	http.Handle("/ws", websocket.Handler(pubSocket.HandleSocket))

	webSocketPort := 8079

	log.Printf("Websocket Connection is Listening on : %v\n", webSocketPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", webSocketPort), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
