package esource

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

var (
	ImpossibleAppendEventOnStream = errors.New("impossible append event on a stream")
)

type StreamName string

type EventStore interface {
	Append(streamName StreamName, events []Event) error
}

type PostgresEventStore struct {
	connection *sql.Conn
}

func (e *PostgresEventStore) Append(streamName StreamName, events []Event) error {
	ctx := context.Background()
	stmt, err := e.connection.PrepareContext(ctx, "")
	defer stmt.Close()

	if err != nil {
		return errors.Wrap(ImpossibleAppendEventOnStream, err.Error())
	}

	_, err = stmt.ExecContext(ctx, "")
	if err != nil {
		return errors.Wrap(ImpossibleAppendEventOnStream, err.Error())
	}

	return nil
}
