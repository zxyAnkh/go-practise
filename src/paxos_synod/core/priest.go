package core

import (
	pb "../protos"
	"fmt"
)

type Priest struct {
	Id        uint32
	Leger     *Leger
	Notes     *[]Note
	Messenger *Messenger
}

var (
	the_Priest Priest // The Priest which this node represented
	the_Decree string // The Decree which is dealing
	// The Flag means whether the_Decree complete consensus.
	// Only complete the_Decree's consensus, the next decree could be handle.
	finished bool = false
)

func InitPriest(id int, nodes []*NodeInfo) error {
	leger, err := InitLeger()
	if err != nil {
		return fmt.Errorf("Init leger error: %v\n", err)
	}
	notes, err := InitNote()
	if err != nil {
		return fmt.Errorf("Init notes error: %v\n", err)
	}
	destinations := make([]NodeInfo, len(nodes)-1)
	i := 0
	for k, v := range nodes {
		if k != id {
			destinations[i], i = *v, i+1
		}
	}
	the_Priest = Priest{
		Id:        uint32(id),
		Leger:     leger,
		Notes:     notes,
		Messenger: NewMessenger(destinations),
	}
	return nil
}

func (p *Priest) dealNewBallotRequest(decree string) {
	if !finished || decree == the_Decree {
		return
	}
	exists := p.Leger.ContainsDecree(decree)
	if exists {
		return
	}
	exists = ContainsNote(*p.Notes, decree)
	if exists {
		return
	}
	var err error
	lastVotes := make([]pb.LastVote, len(p.Messenger.Destination))
	for k, v := range p.Messenger.Destination {
		lastVotes[k], err = p.Messenger.SendPreBallot(v, &pb.NextBallot{
			Id:     1,
			Priest: the_Priest.Id,
		})
		if err != nil {
			fmt.Printf("Can't get message from %s, error: %v\n", v.Ip+":"+v.ServerPort, err)
		}
	}
}

func (p *Priest) dealPreBallot() {

}

func (p *Priest) dealBallot() {

}

func (p *Priest) dealRecordDecree() {

}
