package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	PingDuration = 5 * time.Second
)

type WebSocket struct {
	connection *websocket.Conn
}

func NewWebSocket(connection *websocket.Conn) *WebSocket {
	return &WebSocket{
		connection: connection,
	}
}

func (webSocket *WebSocket) Close() error {
	return webSocket.connection.Close()
}

func (webSocket *WebSocket) Write(message *Message) error {
	return webSocket.connection.WriteMessage(websocket.TextMessage, []byte(message.Content()))
}

func (webSocket *WebSocket) Read() (*Message, error) {
	messageType, content, err := webSocket.connection.ReadMessage()
	if err != nil {
		return &Message{}, err
	}

	return  &Message{messageType: messageType, messageContent: string(content)}, nil
}

func (webSocket *WebSocket) Ping() error {
	return webSocket.connection.WriteMessage(websocket.PingMessage, []byte("Ping"))
}
