package esource

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"sync"
)

var (
	ImpossibleApplyEventChanges = errors.New("impossible apply event changes")
)

type (
	AggregateId  uuid.UUID
	EventApplier interface {
		Apply(events ...Event) error
	}
	EventRecorder interface {
		Events() []Event
		Record(event Event, applier EventApplier) error
	}
	AggregateRoot interface {
		EventApplier
		EventRecorder
		AggregateId() AggregateId
	}
	EventRecorderBase struct {
		events []Event
		mutex  *sync.Mutex
	}
)

func (a *EventRecorderBase) Events() []Event {
	return a.events
}

func (a *EventRecorderBase) Record(event Event, applier EventApplier) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.events = append(a.events, event)
	if err := applier.Apply(event); err != nil {
		return errors.Wrap(ImpossibleApplyEventChanges, err.Error())
	}

	return nil
}
