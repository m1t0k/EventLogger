package types

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

/*
Root type for event registration
*/
type Event struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code       string        `json:"code"`
	Message    string        `json:"message"`
	CreatedAt  time.Time     `json:"createdAt"`
	Priority   string        `json:"priority"`
	Category   string        `json:"category"`
	CustomInfo string        `json:"customInfo"`
}
