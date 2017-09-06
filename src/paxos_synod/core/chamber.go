package core

import (
	pb "../protos"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"net/http"
)

type Chamber struct {
}

func NewChamber() *Chamber {
	return &Chamber{}
}

func InitChamber(ip, serverPort, httpPort string) {
	chamber := NewChamber()
	go chamber.StartServer(ip, serverPort)
	chamber.StartHttpServer(ip, httpPort)
}

/*********************************
*************gRPC Server**********
*********************************/
func (c *Chamber) StartServer(ip, port string) {
	lis, err := net.Listen("tcp", ip+":"+port)
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
	// not deal any ballots util the last ballot progress finish
	if !finished {
		return &pb.LastVote{
			Id:     in.Id,
			MaxId:  0,
			Priest: the_Priest.Id,
		}, nil
	}
	finished = false
	// if already has same ballot id, return default nil struct
	// because of notes' item is more than leger's, so only check notes
	var exists bool = false
	var maxId uint32 = 0
	for _, v := range *the_Priest.Notes {
		if v.Id == in.Id {
			exists = true
		} else if v.Id < in.Id && uint32(v.Priest) == the_Priest.Id && v.Id > maxId {
			maxId = v.Id
		}
	}
	if exists {
		finished = true
		maxId = 0
	}
	r := &pb.LastVote{
		Id:     in.Id,
		MaxId:  maxId,
		Priest: the_Priest.Id,
	}
	fmt.Println(r)
	return r, nil
}

func (c *Chamber) DealBallot(ctx context.Context, in *pb.BeginBallot) (*pb.Voted, error) {
	return &pb.Voted{
		Vote:   false,
		Id:     1,
		Priest: 1,
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

/*********************************
*************HTTP Server**********
**********************************
*the way to produce a new decree*/
func (c *Chamber) StartHttpServer(ip, port string) {
	http.HandleFunc("/synod/"+fmt.Sprintf("%d", the_Priest.Id)+"/", ChamberHttpServer)
	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		fmt.Errorf("Error: %v\n", err)
	}
}

// only deal post request with formatted data
// {"decree":"make this world better."}
func ChamberHttpServer(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqBody := req.Body
		buf, err := ioutil.ReadAll(reqBody)
		if err != nil {
			fmt.Errorf("Read body from request error: %v\n", err)
		}
		decree, err := getDecreeContent(buf)
		if err != nil {
			fmt.Errorf("Get decree from body error: %v\n", err)
		}
		fmt.Println("Decree is :", decree)
		the_Priest.dealNewBallotRequest(decree)
	}
}

// [123 100 101 99 114 101 101 58 ... 125]
func getDecreeContent(buf []byte) (string, error) {
	if string(buf[:8]) == "{decree:" && buf[len(buf)-1] == 125 {
		buf = buf[8:]
		buf = buf[:len(buf)-1]
		return string(buf), nil
	}
	return "", fmt.Errorf("Not formatted data\n")
}
