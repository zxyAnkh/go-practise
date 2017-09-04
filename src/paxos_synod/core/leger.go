package core

import (
	"fmt"
)

type LegerItem struct {
	Id     uint32
	Decree string
}

type Leger struct {
	Items []LegerItem
}

/***********************************
********Leger Item functions********
************************************/

func NewLegerItem(id uint32, decree string) *LegerItem {
	return &LegerItem{
		Id:     id,
		Decree: decree,
	}
}

func (li *LegerItem) equals(item LegerItem) bool {
	return item.Id == li.Id && item.Decree == li.Decree
}

/***********************************
************Leger functions*********
************************************/

// read leger from db
func InitLeger() (*Leger, error) {
	newLeger := &Leger{}
	oldLeger, err := FindAllLegerItem()
	if err != nil {
		return newLeger, err
	}
	newLeger.Items = oldLeger.Items
	return newLeger, nil
}

// add item to leger
// the leger initialize length is 5
// capacity will double if the length equals capacity
func (l *Leger) AddItem(item LegerItem) error {
	if exists, _ := l.ContainsItem(item); exists == true {
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
func (l *Leger) ContainsItem(item LegerItem) (bool, error) {
	for _, v := range l.Items {
		if v.Id == item.Id && v.Decree == item.Decree {
			return true, nil
		}
	}
	return false, nil
}

// judge whether the decree is in leger
func (l *Leger) ContainsDecree(decree string) bool {
	for _, v := range l.Items {
		if decree == v.Decree {
			return true
		}
	}
	return false
}

func (l *Leger) ContainsId(id uint32) bool {
	for _, v := range l.Items {
		if id == v.Id {
			return true
		}
	}
	return false
}
