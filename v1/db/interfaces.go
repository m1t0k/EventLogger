package db

import (
	EventTypes "../types"
)

type iEventDbProvider interface {
	InsertEvent(event EventTypes.Event) bool
	GetEventList() ([]EventTypes.Event, error)
}
