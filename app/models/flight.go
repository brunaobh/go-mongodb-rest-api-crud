package models

import "gopkg.in/mgo.v2/bson"

// Represents a flight
type Flight struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	FlightNum   string        `bson:"flightNum" json:"flightNum"`
	Airline     string        `bson:"airline" json:"airline"`
	Airport     string        `bson:"airport" json:"airport"`
	Status      string        `bson:"status" json:"status"`
	Expected    string        `bson:"expected" json:"expected"`
	Confirmed   string        `bson:"confirmed" json:"confirmed"`
}
