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
	p.Notes, err = initNote()
	if err != nil {
		return
	}
	var count int = len(p.Messenger.Destination)
	lastVotes := make([]*pb.LastVote, count)
	for k, v := range p.Messenger.Destination {
		lastVotes[k], err = p.Messenger.sendPreBallot(v, &pb.NextBallot{
			Id:     id,
			Priest: the_Priest.Id,
		})
		if err != nil {
			fmt.Printf("Can't get message from %s, error: %v\n", v.Ip+":"+v.ServerPort, err)
			count--
		}
	}
	if count != len(p.Messenger.Destination) {
		return
	}
	the_Priest.dealPreBallotResponse(lastVotes, id, decree)
}

// lastVotes: other priest's response; id: this ballot id; decree: this ballot decree(init)
func (p *Priest) dealPreBallotResponse(lastVotes []*pb.LastVote, id uint32, decree string) {
	var maxId uint32 = 0
	for _, v := range lastVotes {
		if v.MaxId != 0 && maxId < v.MaxId && !the_Priest.Leger.containsId(v.MaxId) {
			maxId = v.MaxId
		}
	}
	if maxId != 0 && id > maxId {
		for _, v := range the_Priest.Notes {
			if v.Id == maxId {
				id = v.Id
				decree = v.Decree
				break
			}
		}
	}
	note := Note{
		Id:     id,
		Decree: decree,
		Priest: int(the_Priest.Id),
	}
	var err error
	the_Priest.Notes, err = addNote(the_Priest.Notes, note)
	if err != nil {
		fmt.Println("Add Note error: %v\n", err)
		return
	}
	beginBallot := &pb.BeginBallot{
		Id:     id,
		Decree: decree,
		Priest: the_Priest.Id,
	}
	voteds := make([]*pb.Voted, len(p.Messenger.Destination))
	for k, v := range p.Messenger.Destination {
		voteds[k], err = p.Messenger.sendBallot(v, beginBallot)
		if err != nil {
			fmt.Println("Can't get message from %s, error: %v\n", v.Ip+":"+v.ServerPort, err)
		}
	}
	if len(voteds) != len(p.Messenger.Destination) {
		return
	}
	the_Priest.dealBallotResponse(id, decree)
}

func (p *Priest) dealBallotResponse(id uint32, decree string) {
	err := the_Priest.Leger.addItem(LegerItem{
		Id:     id,
		Decree: decree,
		Priest: the_Priest.Id,
	})
	if err != nil {
		fmt.Println("Insert leger item into db error: %v\n", err)
	}
	success := &pb.Success{
		Id:     id,
		Decree: decree,
		Priest: the_Priest.Id,
	}
	var count int = 0
	for _, v := range p.Messenger.Destination {
		err = p.Messenger.sendRecordDecree(v, success)
		if err != nil {
			fmt.Println("Can't get message from %s, error: %v\n", v.Ip+":"+v.ServerPort, err)
		} else {
			count++
		}
	}
	if count == len(p.Messenger.Destination) {
		fmt.Println("Add a new decree....")
	}
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
