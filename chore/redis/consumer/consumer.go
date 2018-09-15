package consumer

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/loredami/server/pkg/pubsub"
	"log"
	"net/http"
)

func main() {
	rc := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	pubSubA, _ := pubsub.NewPubSub("A", rc)
	pubSubB, _ := pubsub.NewPubSub("B", rc)
	pubSubC, _ := pubsub.NewPubSub("C", rc)
	pubSubs := map[string]*pubsub.RedisPubSub{
		"A": pubSubA,
		"B": pubSubB,
		"C": pubSubC,
	}

	for _, pubSub := range pubSubs {
		go func(pubSub pubsub.RedisPubSub) {
			for {
				fmt.Println("waiting")
				msg, _ := pubSub.Receive()
				fmt.Println(pubSub.Name(), " received: ", msg)
				fmt.Println("loooooping")
			}
		}(*pubSub)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
