package db

import (
	"../../core"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	db           string = "paxos_synod"
	collec_leger string = "leger"
	collec_note  string = "note"
)

func InsertLeger(item core.LegerItem) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_leger)
	err = c.Insert(item)
	if err != nil {
		return fmt.Errorf("Inser leger item error: %v\n", err)
	}
	return nil
}

func FindOneLegerItem() (core.LegerItem, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_leger)
	result := core.LegerItem{}
	err = c.Find(bson.M{"id": 1}).One(&result)
	if err != nil {
		return result, fmt.Errorf("error: %v\n", err)
	}
	return result, nil
}
