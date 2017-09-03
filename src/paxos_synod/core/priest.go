package core

import (
	"fmt"
)

type Priest struct {
	Id    int
	Leger *Leger
	Notes *[]Note
}

var (
	the_Priest Priest
)

func InitPriest(id int) error {
	leger, err := InitLeger()
	if err != nil {
		return fmt.Errorf("Init leger error: %v\n", err)
	}
	notes, err := InitNote()
	if err != nil {
		return fmt.Errorf("Init notes error: %v\n", err)
	}
	the_Priest = Priest{
		Id:    id,
		Leger: leger,
		Notes: notes,
	}
	return nil
}
