package pubsub

import (
	"errors"
	"github.com/go-redis/redis"
)

type Message struct {
	content string
}

type RedisPubSub struct {
	name        string
	redisPubSub *redis.PubSub
	client      *redis.Client
}

type PubSub interface {
	Name() string
	Send(message Message) error
	Receive() (*Message, error)
	Close() error
}

func NewMessage(content string) *Message {
	return &Message{
		content: content,
	}
}

func (redisMessage Message) Content() string {
	return redisMessage.content
}

func NewRedisPubSub(name string, client *redis.Client) (*RedisPubSub, error) {
	redisPubSub := client.Subscribe(name)

	if _, err := redisPubSub.Receive(); err != nil {
		return &RedisPubSub{}, errors.New("impossible open redisPubSub due to: " + err.Error())
	}

	return &RedisPubSub{
		name:        name,
		redisPubSub: redisPubSub,
		client:      client,
	}, nil
}

func (pubSub RedisPubSub) Name() string {
	return pubSub.name
}

func (pubSub RedisPubSub) Send(message Message) error {
	if err := pubSub.client.Publish(pubSub.name, message.Content()).Err(); err != nil {
		return errors.New("impossible send message to redis. Due to: " + err.Error())
	}

	return nil
}

func (pubSub RedisPubSub) Receive() (*Message, error) {
	message, ok := <-pubSub.redisPubSub.Channel()
	if ok != true {
		return &Message{}, errors.New("an error occurred when reading a message")
	}

	return &Message{
		content: message.Payload,
	}, nil
}

func (pubSub RedisPubSub) Close() error {
	return pubSub.redisPubSub.Close()
}
