package pubsub

import (
	"errors"
	"github.com/go-redis/redis"
)

type Message struct {
	content string
}

type PubSub struct {
	name        string
	redisPubSub *redis.PubSub
	client      *redis.Client
}

func NewMessage(content string) *Message {
	return &Message{
		content: content,
	}
}

func (redisMessage Message) Content() string {
	return redisMessage.content
}

func (pubSub PubSub) Name() string {
	return pubSub.name
}

func NewPubSub(name string, client *redis.Client) (*PubSub, error) {
	redisPubSub := client.Subscribe(name)

	if _, err := redisPubSub.Receive(); err != nil {
		return &PubSub{}, errors.New("impossible open redisPubSub due to: " + err.Error())
	}

	return &PubSub{
		name:        name,
		redisPubSub: redisPubSub,
		client:      client,
	}, nil
}

func (pubSub PubSub) Send(message Message) error {
	if err := pubSub.client.Publish(pubSub.name, message.Content()).Err(); err != nil {
		return errors.New("impossible send message to redis. Due to: " + err.Error())
	}

	return nil
}

func (pubSub PubSub) Receive() (*Message, error) {
	msg, ok := <-pubSub.redisPubSub.Channel()
	if ok != true {
		return &Message{}, errors.New("an error occurred when reading a message")
	}

	return &Message{
		content: msg.Payload,
	}, nil
}

func (pubSub PubSub) Close() error {
	return pubSub.redisPubSub.Close()
}
