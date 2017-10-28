package core

import (
	"fmt"
	"ioutil"
	"net"
	"net/http"
	"strconv"
)

const (
	Follower_Status = 1 + iota
	Candidate_Status
	Leader_Status
)

type Server struct {
	Status     int
	RPCConn    *net.UDPConn
	ServerPort int
	RPCPort    int
}

func NewServer(serverPort, rpcPort int) *Server {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(rpcPort))
	if err != nil {
		return &Server{}
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return &Server{}
	}
	server := &Server{
		Status:  Follower_Status,
		RPCConn: conn,
	}
	return server
}

func (server *Server) Start() {
	http.HandlerFunc("/raft/", server.raftHttpServer)
	err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(server.ServerPort), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (server *Server) raftHttpServer(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqBody := req.Body
		buf, err := ioutil.ReadAll(reqBody)
		if err != nil {
			fmt.Printf("Read body from request error: %v\n", err)
		}
		if server.Status == Leader_Status {
			server.deal()
		} else {
			server.forwardToLeader()
		}
	}
}

func (server *Server) deal() {

}

func (server *Server) forwardToLeader() {

}
