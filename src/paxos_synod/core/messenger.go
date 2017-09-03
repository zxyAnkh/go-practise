package core

import (
	"fmt"
	"os"

	pb "../protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Messenger struct {
	Destination []NodeInfo
}

func NewMessager(destinations []NodeInfo) *Messenger {
	return &Messenger{
		Destination: destinations,
	}
}

func InitMessager() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.Request{Str: name})
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}
	fmt.Printf("Greeting: %s\n", r.Str)
}
