package websocket

import (
	"errors"
	"github.com/loredami/server/pkg/pubsub"
	log "github.com/sirupsen/logrus"
	"time"
)

type ClientId string

type Client struct {
	id         ClientId
	hub        *Hub
	webSockets map[*WebSocket]bool
	Register   chan *WebSocket
	Unregister chan *WebSocket
	pubSub     pubsub.PubSub
	logger     *log.Logger
	Write      chan *Message
	Read       chan *Message
}

func NewClient(clientId ClientId, hub *Hub, pubSub pubsub.PubSub, logger *log.Logger) *Client {
	client := Client{
		id:         clientId,
		hub:        hub,
		pubSub:     pubSub,
		webSockets: map[*WebSocket]bool{},
		logger:     logger,
		Register:   make(chan *WebSocket),
		Unregister: make(chan *WebSocket),
		Read:       make(chan *Message),
		Write:      make(chan *Message),
	}
	go client.listen()
	return &client
}

func (client *Client) Id() ClientId {
	return client.id
}

func (client *Client) listen() {
	client.logger.Info("Start listening to the client")
	ticker := time.NewTicker(PingDuration)
	defer ticker.Stop()

	go readingFromPubSub(client)

	for {
		select {
		case webSocket := <-client.Register:
			client.webSockets[webSocket] = true
			go client.readingFromWebSocket(webSocket)
		case webSocket := <-client.Unregister:
			client.removeWebSocket(webSocket)
		case message := <-client.Write:
			for webSocket := range client.webSockets {
				if err := webSocket.Write(message); err != nil {
					client.Unregister <- webSocket
				}
			}
		case message := <-client.Read:
			client.logger.Info("We ignore messages for now :)", message.Content())
		case <-ticker.C:
			for webSocket := range client.webSockets {
				if err := webSocket.Ping(); err != nil {
					client.Unregister <- webSocket
				}
			}
			if len(client.webSockets) == 0 {
				client.pubSub.Close()
				client.hub.Unregister <- client
				return
			}
		}
	}
}

func (client *Client) removeWebSocket(webSocket *WebSocket) error {
	if _, exists := client.webSockets[webSocket]; !exists {
		return errors.New("websocket does not exists")
	}
	delete(client.webSockets, webSocket)
	return webSocket.Close()
}

func (client *Client) readingFromWebSocket(webSocket *WebSocket) {
	for {
		message, err := webSocket.Read()
		if err != nil {
			client.Unregister <- webSocket
			break
		}

		client.Read <- message
	}
}

func readingFromPubSub(client *Client) {
	for {
		message, err := client.pubSub.Receive()
		if err != nil {
			client.hub.Unregister <- client
			return
		}

		client.Write <- &Message{messageType: TextMessage, messageContent: message.Content()}
	}
}
