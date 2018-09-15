package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	PingDuration = 5 * time.Second
)

type WebSocket interface {
	Ping() error
	Read() (*Message, error)
	Write(message *Message) error
	Close() error
}

type GorillaWebSocket struct {
	connection *websocket.Conn
}

func NewGorillaWebSocket(connection *websocket.Conn) *GorillaWebSocket {
	return &GorillaWebSocket{
		connection: connection,
	}
}

func (webSocket *GorillaWebSocket) Close() error {
	return webSocket.connection.Close()
}

func (webSocket *GorillaWebSocket) Write(message *Message) error {
	return webSocket.connection.WriteMessage(websocket.TextMessage, []byte(message.Content()))
}

func (webSocket *GorillaWebSocket) Read() (*Message, error) {
	messageType, content, err := webSocket.connection.ReadMessage()
	if err != nil {
		return &Message{}, err
	}

	return &Message{messageType: messageType, messageContent: string(content)}, nil
}

func (webSocket *GorillaWebSocket) Ping() error {
	return webSocket.connection.WriteMessage(websocket.PingMessage, []byte("Ping"))
}
