package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/loredami/server/pkg/pubsub"
	websocket2 "github.com/loredami/server/pkg/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // add origin check
		},
	}

	hub := websocket2.NewHub(logger)

	go func() {
		for range time.Tick(time.Second * 1) {
			fmt.Println("Size of connected client is ", hub.CountClients())
		}
	}()

	http.HandleFunc("/ws", func(response http.ResponseWriter, request *http.Request) {
		upgradedConnection, err := upgrader.Upgrade(response, request, nil)
		if err != nil {
			fmt.Fprint(response, "impossible upgrade connection. Due to: "+err.Error())
			return
		}
		logger.Info("Connection upgraded")
		webSocket := websocket2.NewGorillaWebSocket(upgradedConnection)
		clientId := websocket2.ClientId("helo") // change in JWT and take userId
		logger.Info("Connecting to redis")
		pubSub, err := pubsub.NewRedisPubSub(string(clientId), redisClient)
		if err != nil {
			fmt.Fprint(response, "impossible subsbribe user to notifications. Due to: "+err.Error())
			return
		}

		if !hub.HasClient(clientId) {
			logger.Info("Client not found in hub")
			client := websocket2.NewClient(clientId, hub, pubSub, logger)
			client.Register <- webSocket
			hub.Register <- client
			return
		}

		logger.Info("Client found in hub")
		client, err := hub.GetClient(clientId)
		if err != nil {
			fmt.Fprint(response, "impossible subsbribe user to notifications. Due to: "+err.Error())
			return
		}
		client.Register <- webSocket
	})

	http.ListenAndServe(":81", nil)
}
