package esource

import (
	"github.com/satori/go.uuid"
	"time"
)

type EventType string

type Event interface {
	Id() uuid.UUID
	AggregateID() AggregateId
	EventType() EventType
	Payload() interface{}
	Timestamp() time.Time
}
