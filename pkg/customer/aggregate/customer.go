package aggregate

import (
	"fmt"

	"github.com/loredami/server/internal/esource"
	customerEvent "github.com/loredami/server/pkg/customer/event"
	"github.com/loredami/server/pkg/customer/valueObject"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

var InvalidCustomerCreatedPayload = errors.New("given customerEvent is not of the expected type")

type Customer struct {
	esource.EventRecorderBase

	id       valueObject.CustomerId
	fullName valueObject.FullName
	email    valueObject.Email
	password valueObject.Password
}

func NewCustomer(
	id valueObject.CustomerId,
	fullName valueObject.FullName,
	email valueObject.Email,
	password valueObject.Password,
) *Customer {

	c := Customer{
		id:       id,
		fullName: fullName,
		email:    email,
		password: password,
	}

	c.Record(customerEvent.NewCustomerCreatedEvent(id, fullName, email, password), &c)

	return &c
}

func (c *Customer) AggregateId() esource.AggregateId {
	return esource.AggregateId(c.id.Value())
}

func (c *Customer) Id() valueObject.CustomerId {
	return c.id
}

func (c *Customer) FullName() valueObject.FullName {
	return c.fullName
}

func (c *Customer) Email() valueObject.Email {
	return c.email
}

func (c *Customer) Password() valueObject.Password {
	return c.password
}

func (c *Customer) Apply(events ...esource.Event) error {
	for _, event := range events {
		switch eventType := event.(type) {
		case customerEvent.CustomerCreatedEvent:
			if err := c.applyCustomerCreatedEvent(eventType); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Customer) applyCustomerCreatedEvent(event customerEvent.CustomerCreatedEvent) error {
	c.id = *valueObject.NewCustomerId(uuid.UUID(event.AggregateID()))
	payload, ok := event.Payload().(customerEvent.CustomerCreatedPayload)
	if !ok {
		return errors.Wrap(InvalidCustomerCreatedPayload, fmt.Sprint("customerEvent with id: ", event.Id()))
	}
	c.fullName = payload.FullName()
	c.email = payload.Email()
	c.password = payload.Password()

	return nil
}
