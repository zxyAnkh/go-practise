package core

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
)

const (
	db_server     string = "localhost:6379"
	db_leger      string = "paxos.synod.leger"
	db_note       string = "paxos.synod.note"
	db_leger_size string = "paxos.synod.leger.size"
	db_note_size  string = "paxos.synod.note.size"
)

// TODO: marshal/unmarshal key

var redis_client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     db_server,
	Password: "",
	DB:       0,
})

/*********************
*********Leger********
*********************/
func insertLegerItem(item LegerItem) error {
	val, err := json.Marshal(item)
	if err != nil {
		return err
	}
	err = redis_client.Set(db_leger+"."+strconv.Itoa(int(item.Id)), string(val), 0).Err()
	if err != nil {
		return err
	}
	return updateLegerSize()
}

func updateLegerSize() error {
	val, err := redis_client.Get(db_leger_size).Result()
	if err == redis.Nil {
		err = redis_client.Set(db_leger_size, "0", 0).Err()
		return err
	} else if err != nil {
		return nil
	}
	size, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	err = redis_client.Set(db_leger_size, strconv.Itoa(size+1), 0).Err()
	return err
}

func findOneLegerItemById(id uint) (LegerItem, error) {
	val, err := redis_client.Get(db_leger + "." + strconv.Itoa(int(id))).Result()
	if err == redis.Nil {
		return LegerItem{}, nil
	} else if err != nil {
		return LegerItem{}, err
	}
	var item LegerItem
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func findAllLegerItem() (Leger, error) {
	val, err := redis_client.Get(db_leger_size).Result()
	if err == redis.Nil {
		return Leger{}, nil
	} else if err != nil {
		return Leger{}, err
	}
	size, err := strconv.Atoi(val)
	if err != nil {
		return Leger{}, err
	}
	leger := Leger{
		Items: make([]LegerItem, size),
	}
	for i := 1; i <= size; i++ {
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
	val, err := json.Marshal(note)
	if err != nil {
		return err
	}
	err = redis_client.Set(db_note+"."+strconv.Itoa(int(note.Id)), string(val), 0).Err()
	if err != nil {
		return err
	}
	return updateNoteSize()
}

func updateNoteSize() error {
	val, err := redis_client.Get(db_note_size).Result()
	if err == redis.Nil {
		err = redis_client.Set(db_note_size, "0", 0).Err()
		return err
	} else if err != nil {
		return nil
	}
	size, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	err = redis_client.Set(db_note_size, strconv.Itoa(size+1), 0).Err()
	return err
}

func deleteNoteById(id uint) error {
	err := redis_client.Del(db_note + "." + strconv.Itoa(int(id))).Err()
	return err
}

func updateNote(oldNote, newNote Note) error {
	val, err := json.Marshal(newNote)
	if err != nil {
		return err
	}
	err = redis_client.Set(db_note+"."+strconv.Itoa(int(oldNote.Id)), string(val), 0).Err()
	return err
}

func findOneNoteById(id uint) (Note, error) {
	val, err := redis_client.Get(db_note + "." + strconv.Itoa(int(id))).Result()
	if err == redis.Nil {
		return Note{}, nil
	} else if err != nil {
		return Note{}, err
	}
	var note Note
	err = json.Unmarshal([]byte(val), &note)
	if err != nil {
		return note, err
	}
	return note, nil
}

func findAllNote() ([]Note, error) {
	val, err := redis_client.Get(db_note_size).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	size, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	notes := make([]Note, size)
	for i := 1; i <= size; i++ {
		note, err := findOneNoteById(uint(i))
		if err != nil {
			notes[i-1] = note
		}
	}
	return notes, nil
}
