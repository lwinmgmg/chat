package socket

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lwinmgmg/chat/exceptions"
	"golang.org/x/net/websocket"
)

func waitForMesgSize(buffer []byte, ws *websocket.Conn) (int, error) {
	n, err := ws.Read(buffer)
	if err != nil {
		return 0, err
	}
	sizeStr := string(buffer[:n])
	sizeStr = strings.Replace(sizeStr, " ", "", -1)
	return strconv.Atoi(sizeStr)
}

func (socketHandler *SocketHandler) ReadMesg(ws *websocket.Conn) ([]byte, error) {
	buff := make([]byte, 5)
	size, err := waitForMesgSize(buff, ws)
	if err != nil {
		return nil, errors.Join(err, exceptions.ErrReadSize)
	}
	mesgBytes := make([]byte, size)
	readedSize, err := ws.Read(mesgBytes)
	if err != nil {
		return nil, errors.Join(err, exceptions.ErrReadMesg)
	}
	return mesgBytes[:readedSize], nil
}
