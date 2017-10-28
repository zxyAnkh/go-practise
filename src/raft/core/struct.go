package core

// RPC
type AppendEntries struct {
	Term         uint32  `json: "term"`
	LeaderId     uint32  `json: "leaderId"`
	PrevLogIndex uint32  `json: "prevLogIndex"`
	entries      []Entry `json: "entries"`
	LeaderCommit uint32  `json: "leaderCommit"`
}
type AppendEntriesResponse struct {
	Term    uint32 `json: "term"`
	Success bool   `json: "success"`
}
type RequestVote struct {
	Term         uint32 `json: "term"`
	CandidateId  uint32 `json: "candidateId"`
	LastLogIndex uint32 `json: "lastLogIndex"`
	LastLogTerm  uint32 `json: "lastLogTerm"`
}
type ReqeustVoteResponse struct {
	Term        uint32 `json: "term"`
	VoteGranted bool   `json: "voteGranted"`
}
type Entry struct {
	Index   uint32 `json: "index"`
	Term    uint32 `json: "term"`
	Command string `json: "command"`
}

// Role
type Leader struct {
	NextIndex  []uint32
	MatchIndex []uint32
}
type Candidate struct {
}
type Follower struct {
	CurrentTerm uint32
	VotedFor    uint32
	Log         []Entry
	CommitIndex uint32
	LastApplied uint32
}
