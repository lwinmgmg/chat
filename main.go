package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/chat/controllers"
	"github.com/lwinmgmg/chat/socket"
	"golang.org/x/net/websocket"
)

// This example demonstrates a trivial echo server.
func main() {
	// Web Socket
	pubSocket := socket.SocketHandler{
		Auth:    false,
		UserMap: make(map[string]*socket.UserInfo, 100),
		ConnAge: time.Hour,
	}
	pubSocket.Init()
	defer pubSocket.Close()
	http.Handle("/ws", websocket.Handler(pubSocket.HandleSocket))

	webSocketPort := 8079

	log.Printf("Websocket Connection is Listening on : %v\n", webSocketPort)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(fmt.Sprintf(":%v", webSocketPort), nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	app := gin.Default()
	app.Use(cors.Default())
	controllers.DefineRoutes(app)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Run("localhost:8078"); err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}
