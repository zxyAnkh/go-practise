package main

import (
	pb "../protobuf"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":12345"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Str: "hello"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
