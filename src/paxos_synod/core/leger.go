package core

import (
	"fmt"
)

type LegerItem struct {
	Id     uint32
	Decree string
	Priest uint32
}

type Leger struct {
	Items []LegerItem
}

/***********************************
********Leger Item functions********
************************************/

func newLegerItem(id, priest uint32, decree string) *LegerItem {
	return &LegerItem{
		Id:     id,
		Decree: decree,
		Priest: priest,
	}
}

func (li *LegerItem) equals(item LegerItem) bool {
	return item.Id == li.Id && item.Decree == li.Decree && item.Priest == li.Priest
}

/***********************************
************Leger functions*********
************************************/

// read leger from db
func initLeger() (*Leger, error) {
	newLeger := &Leger{}
	oldLeger, err := findAllLegerItem()
	if err != nil {
		return newLeger, err
	}
	newLeger.Items = oldLeger.Items
	return newLeger, nil
}

// add item to leger
// the leger initialize length is 5
// capacity will double if the length equals capacity
func (l *Leger) addItem(item LegerItem) error {
	if exists, _ := l.containsItem(item); exists == true {
		return fmt.Errorf("Item %v is exists.\n", item)
	}
	length := len(l.Items)
	if length == 0 {
		l.Items = make([]LegerItem, 5)
	}
	capacity := cap(l.Items)
	if capacity == length {
		newItems := make([]LegerItem, capacity*2)
		copy(newItems, l.Items)
		l.Items = newItems
	}
	l.Items[length+1] = item
	return nil
}

// judge whether the item is in leger
func (l *Leger) containsItem(item LegerItem) (bool, error) {
	for _, v := range l.Items {
		if v.Id == item.Id && v.Decree == item.Decree && v.Priest == item.Priest {
			return true, nil
		}
	}
	return false, nil
}

// judge whether the decree is in leger
func (l *Leger) containsDecree(decree string) bool {
	for _, v := range l.Items {
		if decree == v.Decree {
			return true
		}
	}
	return false
}

func (l *Leger) containsId(id uint32) bool {
	for _, v := range l.Items {
		if id == v.Id {
			return true
		}
	}
	return false
}
