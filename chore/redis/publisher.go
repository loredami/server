package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/loredami/server/pkg/pubsub"
	"math/rand"
	"time"
)

func main() {

	rc := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	pubSubA, _ := pubsub.NewPubSub("A", rc)
	pubSubB, _ := pubsub.NewPubSub("B", rc)
	pubSubC, _ := pubsub.NewPubSub("C", rc)
	pubSubs := map[string]*pubsub.PubSub{
		"A": pubSubA,
		"B": pubSubB,
		"C": pubSubC,
	}

	for {
		for _, pubSub := range pubSubs {
			fmt.Println("sending")
			msg := fmt.Sprint(rand.Int())
			fmt.Println("Error: ", pubSub.Send(*pubsub.NewMessage(msg)))
			fmt.Println("Sent ", msg, " to : ", pubSub.Name())
		}
		time.Sleep(10 * time.Millisecond)
	}
}
