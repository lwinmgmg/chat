package models

type SocketType string

const (
	CHAT SocketType = "chat"
)

type SocketData struct {
	SocketType SocketType     `json:"socket_type,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}
