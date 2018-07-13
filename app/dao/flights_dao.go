package dao

import (
	"log"

	. "github.com/user/app/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FlightsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "flights"
)

// Establish a connection to database
func (m *FlightsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of flights
func (m *FlightsDAO) FindAll() ([]Flight, error) {
	var flights []Flight
	err := db.C(COLLECTION).Find(bson.M{}).All(&flights)
	return flights, err
}

// Find a flight by its id
func (m *FlightsDAO) FindById(id string) (Flight, error) {
	var flight Flight
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&flight)
	return flight, err
}

// Insert a flight into database
func (m *FlightsDAO) Insert(flight Flight) error {
	err := db.C(COLLECTION).Insert(&flight)
	return err
}

// Delete an existing flight
func (m *FlightsDAO) Delete(flight Flight) error {
	err := db.C(COLLECTION).Remove(&flight)
	return err
}

// Update an existing flight
func (m *FlightsDAO) Update(flight Flight) error {
	err := db.C(COLLECTION).UpdateId(flight.ID, &flight)
	return err
}
