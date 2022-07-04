package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Bus defines the expected behaviour from an event bus.
type Bus interface {
	// Publish is the methos used to publish new events.
	Publish(context.Context, []Event) error

	//Suscribe is the method used to subscribe new event handlers.
	Subscribe(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=eventmocks --output=eventmocks --name=Bus

type Handler interface {
	Handle(context.Context, Event) error
}

// Type represents a domain event type.
type Type string

// Event represents a domain command.
type Event interface {
	ID() string
	AggregateID() string
	OccurendOn() time.Time
	Type() Type
}

type BaseEvent struct {
	eventID     string
	aggregateID string
	occurendOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		eventID:     uuid.New().String(),
		aggregateID: aggregateID,
		occurendOn:  time.Now(),
	}
}

func (b BaseEvent) ID() string {
	return b.eventID
}

func (b BaseEvent) OccurendOn() time.Time {
	return b.occurendOn
}

func (b BaseEvent) AggregateID() string {
	return b.aggregateID
}
