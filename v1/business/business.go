package business

import (
	EventDb "../db"
	EventTypes "../types"
	"time"
)

/*
EventProvider type
*/
type EventProvider struct {
	dbp EventDb.EventDbProvider
}

/*
insert event into mongo db
*/
func (ep EventProvider) InsertEvent(event EventTypes.Event) bool {
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now().UTC()
	}
	return ep.dbp.DbInsertEvent(event)
}

/*
/get full event list
*/
func (ep EventProvider) GetEventList() ([]EventTypes.Event, error) {
	return ep.dbp.DbGetEventList()
}

/*
get event by id
*/
func (ep EventProvider) GetEvent(id string) (EventTypes.Event, error) {
	return ep.dbp.DbGetEvent(id)
}
