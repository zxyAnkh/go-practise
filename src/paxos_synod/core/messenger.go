package core

import (
	"fmt"

	pb "../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Messenger struct {
	Destination []NodeInfo
}

func NewMessenger(destinations []NodeInfo) *Messenger {
	return &Messenger{
		Destination: destinations,
	}
}

func (m *Messenger) SendPreBallot(dest NodeInfo, nextBallot *pb.NextBallot) (*pb.LastVote, error) {
	r := &pb.LastVote{}
	conn, err := grpc.Dial(dest.Ip+":"+dest.ServerPort, grpc.WithInsecure())
	if err != nil {
		return r, fmt.Errorf("Connect to %s error: %v\n", dest.Ip+":"+dest.ServerPort, err)
	}
	defer conn.Close()
	c := pb.NewPaxosClient(conn)

	r, err = c.DealPreBallot(context.Background(), nextBallot)
	if err != nil {
		return r, fmt.Errorf("Could not greet: %v\n", err)
	}
	fmt.Println(r)
	return r, nil
}

func (m *Messenger) SendBallot(dest NodeInfo, beginBallot *pb.BeginBallot) (*pb.Voted, error) {
	r := &pb.Voted{}
	conn, err := grpc.Dial(dest.Ip+":"+dest.ServerPort, grpc.WithInsecure())
	if err != nil {
		return r, fmt.Errorf("Connect to %s error: %v\n", dest.Ip+":"+dest.ServerPort, err)
	}
	defer conn.Close()
	c := pb.NewPaxosClient(conn)

	r, err = c.DealBallot(context.Background(), beginBallot)
	if err != nil {
		return r, fmt.Errorf("Could not greet: %v\n", err)
	}
	fmt.Println(r)
	return r, nil
}

func (m *Messenger) SendRecordDecree(dest NodeInfo, success *pb.Success) error {
	conn, err := grpc.Dial(dest.Ip+":"+dest.ServerPort, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("Connect to %s error: %v\n", dest.Ip+":"+dest.ServerPort, err)
	}
	defer conn.Close()
	c := pb.NewPaxosClient(conn)

	_, err = c.RecordDecree(context.Background(), success)
	if err != nil {
		return fmt.Errorf("Could not greet: %v\n", err)
	}
	return nil
}
