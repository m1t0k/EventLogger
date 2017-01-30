package db

import (
	AppConfig "../config"
	EventTypes "../types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
EventDBProvider type
*/
type EventDbProvider struct{}

func dbConnect() (*mgo.Session, error) {
	return mgo.Dial(AppConfig.AppConfiguration.DbServer)
}

func getEventsCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(AppConfig.AppConfiguration.DbName).C("events")
}

/*
insert event into MongoDB database
*/
func (edp EventDbProvider) DbInsertEvent(event EventTypes.Event) bool {
	session, errConn := dbConnect()
	defer session.Close()
	if errConn != nil {
		return false
	}
	var eventsCollection = getEventsCollection(session)
	err := eventsCollection.Insert(event)
	return err == nil
}

/*
 Get event by id.
*/
func (edp EventDbProvider) DbGetEvent(id string) (EventTypes.Event, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return EventTypes.Event{}, errConn
	}
	defer session.Close()

	var event EventTypes.Event
	var eventsCollection = getEventsCollection(session)
	err := eventsCollection.FindId(bson.ObjectIdHex(id)).One(&event)
	return event, err
}

/*
Delete event id
*/
func (edp EventDbProvider) DbDeleteEvent(id string) (bool, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return false, errConn
	}
	defer session.Close()

	var eventsCollection = getEventsCollection(session)
	err := eventsCollection.RemoveId(bson.ObjectIdHex(id))
	return err == nil, err
}

/*
Get full event list
*/
func (edp EventDbProvider) DbGetEventList() ([]EventTypes.Event, error) {
	session, errConn := dbConnect()
	if errConn != nil {
		return nil, errConn
	}
	defer session.Close()

	var events []EventTypes.Event
	var eventsCollection = getEventsCollection(session)
	err := eventsCollection.Find(bson.M{}).All(&events)
	return events, err
}
