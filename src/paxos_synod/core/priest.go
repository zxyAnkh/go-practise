package core

import (
	pb "../protos"
	"fmt"
)

type Priest struct {
	Id        uint32
	Leger     *Leger
	Notes     []Note
	Messenger *Messenger
}

var (
	the_Priest   Priest         // The Priest which this node represented
	the_Decree   string         // The Decree which is dealing
	is_President bool   = false // default president is the first node
)

func InitPriest(id int, nodes []NodeInfo) error {
	if id == 1 {
		is_President = true
	}
	leger, err := initLeger()
	if err != nil {
		return fmt.Errorf("Init leger error: %v\n", err)
	}
	notes, err := initNote()
	if err != nil {
		return fmt.Errorf("Init notes error: %v\n", err)
	}
	destinations := make([]NodeInfo, len(nodes)-1)
	i := 0
	for k, v := range nodes {
		if k+1 != id {
			destinations[i] = v
			i++
		}
	}
	the_Priest = Priest{
		Id:        uint32(id),
		Leger:     leger,
		Notes:     notes,
		Messenger: newMessenger(destinations),
	}
	return nil
}

func (p *Priest) dealNewBallotRequest(decree string) {
	if decree == the_Decree {
		return
	}
	exists := p.Leger.containsDecree(decree)
	if exists {
		return
	}
	exists = containsNote(p.Notes, decree)
	if exists {
		return
	}
	var err error
	var id uint32 = the_Priest.genreateBallotId()
	err = insertNote(Note{
		Id:     id,
		Decree: decree,
		Priest: int(the_Priest.Id),
	})
	p.Notes, err = initNote()
	if err != nil {
		return
	}
	lastVotes := make([]*pb.LastVote, len(p.Messenger.Destination))
	for k, v := range p.Messenger.Destination {
		lastVotes[k], err = p.Messenger.sendPreBallot(v, &pb.NextBallot{
			Id:     id,
			Priest: the_Priest.Id,
		})
		if err != nil {
			fmt.Printf("Can't get message from %s, error: %v\n", v.Ip+":"+v.ServerPort, err)
		}
	}
	if len(lastVotes) != len(p.Messenger.Destination) {
		return
	}
	the_Priest.dealPreBallot(id, decree)
}

func (p *Priest) dealPreBallot(id uint32, decree string) {

}

func (p *Priest) dealBallot() {

}

func (p *Priest) dealRecordDecree() {

}

func (p *Priest) genreateBallotId() uint32 {
	var maxId uint32 = 0
	for _, v := range p.Leger.Items {
		if v.Id > maxId {
			maxId = v.Id
		}
	}
	return maxId + 1
}
