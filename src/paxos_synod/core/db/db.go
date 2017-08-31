package db

import (
	"../../core"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	db_server    string = "localhost"
	db           string = "paxos_synod"
	collec_leger string = "leger"
	collec_note  string = "note"
)

/*********************
*********Leger********
*********************/
func InsertLegerItem(item core.LegerItem) error {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_leger)
	err = c.Insert(item)
	if err != nil {
		return fmt.Errorf("Insert leger item error: %v\n", err)
	}
	return nil
}

func FindOneLegerItemById(id uint) (core.LegerItem, error) {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_leger)
	result := core.LegerItem{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return result, fmt.Errorf("Find one leger item by id error: %v\n", err)
	}
	return result, nil
}

func FindAllLegerItem() (core.Leger, error) {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_leger)
	leger := core.Leger{}
	result := []core.LegerItem{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		return leger, fmt.Errorf("Find all leger item error: %v\n", err)
	}
	leger.Items = result
	return leger, nil
}

/*********************
*********Note*********
*********************/
func InsertNote(note core.Note) error {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_note)
	err = c.Insert(note)
	if err != nil {
		return fmt.Errorf("Insert note error: %v\n", err)
	}
	return nil
}

func DeleteNoteById(id uint) error {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_note)
	err = c.Remove(bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("Delete note by id error: %v\n", err)
	}
	return nil
}

func UpdateNote(note core.Note) error {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_note)
	err = c.Update(bson.M{"id": note.Id, "decree": note.Decree, "priest": note.Priest})
	if err != nil {
		return fmt.Errorf("Update note error: %v\n", err)
	}
	return nil
}

func FindOneNoteById(id uint) (core.Note, error) {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_note)
	result := core.Note{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return result, fmt.Errorf("Find one note by id error: %v\n", err)
	}
	return result, nil
}

func FindAllNote() ([]core.Note, error) {
	session, err := mgo.Dial(db_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collec_note)
	result := []core.Note{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, fmt.Errorf("Find all note error: %v\n", err)
	}
	return result, nil
}
