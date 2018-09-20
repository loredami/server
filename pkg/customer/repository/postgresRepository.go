package repository

import (
	"github.com/loredami/server/internal/esource"
	"github.com/loredami/server/pkg/customer/aggregate"
	"github.com/pkg/errors"
)

type PostgresCustomerRepository struct {
	eventStore esource.EventStore
}

func (r *PostgresCustomerRepository) Add(customer *aggregate.Customer) error {

	if err := r.eventStore.Append(customer.Events(), esource.StreamName("customer")); err != nil {
		return errors.Wrap(ImpossibleAddCustomer, err.Error())
	}

	return nil
}
