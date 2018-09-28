package repository

import (
	"github.com/loredami/server/pkg/customer/aggregate"
	"github.com/pkg/errors"
)

var (
	ImpossibleAddCustomer = errors.New("impossible to add the customer")
)

type CustomerRepository interface {
	Add(customer *aggregate.Customer) error
}
