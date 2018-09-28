package valueObject

import "github.com/satori/go.uuid"

type CustomerId struct {
	value uuid.UUID
}

func NewCustomerId(uuid uuid.UUID) *CustomerId {
	return &CustomerId{
		value: uuid,
	}
}

func (c CustomerId) Value() uuid.UUID {
	return c.value
}

func (c CustomerId) String() string {
	return c.value.String()
}
