package event

import (
	"time"

	"github.com/loredami/server/internal/esource"
	"github.com/loredami/server/pkg/customer/valueObject"
	"github.com/satori/go.uuid"
)

const (
	CustomerCreatesEvent = esource.EventType("customer:created")
)

type CustomerCreatedPayload struct {
	fullName valueObject.FullName
	email    valueObject.Email
	password valueObject.Password
}

type CustomerCreatedEvent struct {
	id          uuid.UUID
	aggregateID esource.AggregateId
	eventType   esource.EventType
	payload     CustomerCreatedPayload
	timestamp   time.Time
}

func NewCustomerCreatedEvent(
	id valueObject.CustomerId,
	fullName valueObject.FullName,
	email valueObject.Email,
	password valueObject.Password,
) *CustomerCreatedEvent {
	return &CustomerCreatedEvent{
		id:          uuid.NewV4(),
		aggregateID: esource.AggregateId(id.Value()),
		eventType:   CustomerCreatesEvent,
		payload:     *NewCustomerCreatedPayload(fullName, email, password),
		timestamp:   time.Now(),
	}
}

func NewCustomerCreatedPayload(
	fullName valueObject.FullName,
	email valueObject.Email,
	password valueObject.Password,
) *CustomerCreatedPayload {
	return &CustomerCreatedPayload{
		fullName: fullName,
		email:    email,
		password: password,
	}
}

func (e CustomerCreatedEvent) Id() uuid.UUID {
	return e.id
}

func (e CustomerCreatedEvent) AggregateID() esource.AggregateId {
	return e.aggregateID
}

func (e CustomerCreatedEvent) EventType() esource.EventType {
	return e.eventType
}

func (e CustomerCreatedEvent) Payload() interface{} {
	return e.payload
}

func (e CustomerCreatedEvent) Timestamp() time.Time {
	return e.timestamp
}

func (p CustomerCreatedPayload) FullName() valueObject.FullName {
	return p.fullName
}

func (p CustomerCreatedPayload) Email() valueObject.Email {
	return p.email
}

func (p CustomerCreatedPayload) Password() valueObject.Password {
	return p.password
}
