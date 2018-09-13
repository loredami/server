package websocket

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Hub struct {
	clients    *map[ClientId]*Client
	Register   chan *Client
	Unregister chan *Client
	logger     *log.Logger
}

func NewHub(logger *log.Logger) *Hub {
	hub := Hub{
		clients:    &map[ClientId]*Client{},
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		logger:     logger,
	}

	go hub.run()
	return &hub
}

func (hub *Hub) run() {
	hub.logger.Info("Starting hub")
	for {
		select {
		case client := <-hub.Register:
			hub.logger.Info("Register client")
			(*hub.clients)[client.id] = client
		case client := <-hub.Unregister:
			hub.logger.Info("Unregister client")
			delete(*hub.clients, client.Id())
		}
	}
}

// REMOVE THIS IS TEMPORARY
func (hub *Hub) CountClients() int {
	return len(*hub.clients)
}

func (hub *Hub) HasClient(clientId ClientId) bool {
	_, exists := (*hub.clients)[clientId]

	return exists
}

func (hub *Hub) GetClient(clientId ClientId) (*Client, error) {

	if !hub.HasClient(clientId) {
		return &Client{}, errors.New(fmt.Sprint("client not found with the clientId: " + clientId))
	}
	return (*hub.clients)[clientId], nil
}
