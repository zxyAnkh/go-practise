package core

import (
	"fmt"
)

type LegerItem struct {
	Id     uint
	Decree string
}

type Leger struct {
	Items []LegerItem
}

/***********************************
********Leger Item functions********
************************************/
func NewLegerItem(id uint, decree string) *LegerItem {
	return &LegerItem{
		Id:     id,
		Decree: decree,
	}
}

/***********************************
************Leger functions*********
************************************/
func NewLeger(items []LegerItem) *Leger {
	return &Leger{
		Items: items,
	}
}

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

func (l *Leger) ContainsItem(item LegerItem) (bool, error) {
	for _, v := range l.Items {
		if v.Id == item.Id && v.Decree == item.Decree {
			return true, nil
		}
	}
	return false, nil
}
