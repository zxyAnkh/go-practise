package core

import (
	"encoding/json"
	"net"
	"os"
)

type RaftRPC interface {
	AppendEntriesRPC(appendEntries AppendEntries)
	ResponseAppendEntriesRPC(appendEntriesRespon AppendEntriesResponse)
	RequestVoteRPC(requestVote RequestVote)
	ResponseRequestVoteRPC(requestVoteResposne ReqeustVoteResponse)
}

func (server *Server) AppendEntriesRPC(appendEntries AppendEntries) {

}

func (server *Server) ResponseAppendEntriesRPC(appendEntriesRespon AppendEntriesResponse) {

}

func (server *Server) RequestVoteRPC(requestVote RequestVote) {

}

func (server *Server) ResponseRequestVoteRPC(requestVoteResposne ReqeustVoteResponse) {

}
