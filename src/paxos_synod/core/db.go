package core

import (
	"github.com/go-redis/redis"
	"strconv"
)

const (
	db_server string = "localhost:6379"
	db_leger  string = "paxos.synod.leger"
	db_note   string = "paxos.synod.note"
)

var redis_client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     db_server,
	Password: "",
	DB:       0,
})

/*********************
*********Leger********
*********************/
func insertLegerItem(item LegerItem) error {
	err := redis_client.Set(db_leger+"."+strconv.Itoa(int(item.Id)), "", 0).Err()
	return err
}

func findOneLegerItemById(id uint) (LegerItem, error) {
	_, err := redis_client.Get(db_leger).Result()
	if err == redis.Nil {
		return LegerItem{}, nil
	} else if err != nil {
		return LegerItem{}, err
	}
	return LegerItem{}, nil
}

func findAllLegerItem() (Leger, error) {
	val, err := redis_client.Get(db_leger).Result()
	if err == redis.Nil {
		return Leger{}, nil
	} else if err != nil {
		return Leger{}, err
	}
	count, err := strconv.Atoi(val)
	if err != nil {
		return Leger{}, err
	}
	leger := Leger{
		Items: make([]LegerItem, count),
	}
	for i := 1; i <= count; i++ {
		item, err := findOneLegerItemById(uint(i))
		if err != nil {
			leger.Items[i-1] = item
		}
	}
	return leger, nil
}

/*********************
*********Note*********
*********************/
func insertNote(note Note) error {
	err := redis_client.Set(db_note, "", 0).Err()
	return err
}

func deleteNoteById(id uint) error {
	err := redis_client.Del("").Err()
	return err
}

func updateNote(oldNote, newNote Note) error {
	err := redis_client.Set(db_note, "", 0).Err()
	return err
}

func findOneNoteById(id uint) (Note, error) {
	_, err := redis_client.Get(db_note).Result()
	if err == redis.Nil {
		return Note{}, nil
	} else if err != nil {
		return Note{}, err
	}
	return Note{}, nil
}

func findAllNote() ([]Note, error) {
	val, err := redis_client.Get(db_leger).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	count, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	notes := make([]Note, count)
	for i := 1; i <= count; i++ {
		note, err := findOneNoteById(uint(i))
		if err != nil {
			notes[i-1] = note
		}
	}
	return notes, nil
}
