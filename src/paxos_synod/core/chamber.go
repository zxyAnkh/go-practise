package core

import (
	pb "../protos"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Chamber struct {
}

func NewChamber() *Chamber {
	return &Chamber{}
}

func (c *Chamber) StartServer(ip, port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Errorf("Start server error: %v\n", err)
	}
	server := grpc.NewServer()
	pb.RegisterPaxosServer(server, c)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		fmt.Errorf("error: %v\n", err)
	}
}

func (c *Chamber) DealPreBallot(ctx context.Context, in *pb.NextBallot) (*pb.LastVote, error) {
	return &pb.LastVote{
		Id:    1,
		MinId: 1,
	}, nil
}

func (c *Chamber) DealBallot(ctx context.Context, in *pb.BeginBallot) (*pb.Voted, error) {
	return &pb.Voted{
		Vote:   false,
		Id:     1,
		Priest: "1",
	}, nil
}

func (c *Chamber) RecordDecree(ctx context.Context, in *pb.Success) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (c *Chamber) Synchronize(ctx context.Context, in *pb.Leger) (*pb.Leger, error) {
	return &pb.Leger{
		Values: make(map[uint32]string, 0),
	}, nil
}
