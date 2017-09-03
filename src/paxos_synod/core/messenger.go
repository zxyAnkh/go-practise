package core

import (
	"fmt"

	pb "../protobuf"
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

func (m *Messenger) SendPreBallot(dest NodeInfo, id uint) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(dest.Ip+":"+dest.ServerPort, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Connect to %s error: %v\n", dest.Ip+":"+dest.ServerPort, err)
	}
	defer conn.Close()
	c := pb.NewPaxosClient(conn)

	r, err := c.DealPreBallot(context.Background(), &pb.NextBallot{Id: uint(id)})
	if err != nil {
		fmt.Printf("Could not greet: %v\n", err)
	}
	fmt.Println(r)
}
